package artillary

import (
	"context"
	"sync"
	"time"

	"github.com/kernelpanic77/hellfire/common"
	"github.com/kernelpanic77/hellfire/internal/client"
	internal "github.com/kernelpanic77/hellfire/internal/client"
	"github.com/kernelpanic77/hellfire/internal/metrics"
)

// contains the general logic of creation of different workers
// ramp up policies etc

// shared iterations - eg: total 100 iterations shared between some VUs - simply create a pool of workers and add number of iterations to job pool
// (number of VUs to run concurrently)

// per vu iterations - each worker should run specific number of iterations - spawn all the workers together withy the same pool func

// constant VUs - for fixed duration - create some Vus at the beginning and then run them for a fixed duration

// ramping VUs up or down based on stages start with specific number of VUs - pool function should run for infinite time - if need ramp up add more, if need to remove kill some worker

// constant-arrival-rate - increase the number of iterations to be executed every timeUNit for a fixed duration with a starting number of VUs and and maxVUs - start with some VUs keep adding until we reach the required number of iterations

// ramping-arrival-rate - ramping same thing with stages -

// lets consider the iteration to be one execution by a worker

// Requirements From a pool

// should report correct number of iterations achievd per timeunit
// should support addition of new workers
// killing a certain number of workers abruptly with grace period
//

// case
// case 1 number of iterations less than the required - assuming (iters/workers) - avg iters per worker -
// like y iters are need in x time and currently w workers achieved z
// then y * w / z workers needed, swiftly increase so many workers in this time frame
//

// Worker is essentially an abstraction over the pool function
// it can run for a fixed number of iterations
// run for constant time
// run until it is killed like performing iterations constantly

type PoolFunc func(interface{})

type Worker struct {
	worker_id             int
	target_iterations     int                          // number of iterations worker must complete
	target_duration       int                          // target duration for the worker
	cutoff_duration       int                          // cutoff the worker after specific time
	iteration_func        common.Task                  // function which worker is supposed to run
	worker_ctx            context.Context              // context for each
	worker_cancel         context.CancelFunc           // cancel function for the context
	worker_ctx_timeout    context.Context              // context for each
	worker_cancel_timeout context.CancelFunc           // cancel function for the context
	kill_worker           chan (bool)                  // kill worker channel
	samples_chan          chan metrics.SampleContainer // dedicated channel for sending metrics
	strategy              Strategy                     // Strategy in workers
	worker_wg             *sync.WaitGroup              // wait group for the pool
}

func NewWorker(id int, iteration_func common.Task, worker_ctx context.Context, samples_chan chan metrics.SampleContainer, wg *sync.WaitGroup) Worker {
	return Worker{
		worker_id:      id,
		iteration_func: iteration_func,
		worker_ctx:     worker_ctx,
		samples_chan:   samples_chan,
		worker_wg:      wg,
	}
}

func (w *Worker) run_iteration() metrics.SampleContainer {
	// create a client for the iterations function and fetch metrics from the client
	// currently lets stick to the http client
	http_client := &internal.Client{}
	w.iteration_func(&client.T{}, http_client)
	metrics := http_client.CollectMetrics()
	return metrics
}

// Run for Iterations
func RunForIterations() PoolFunc {
	return func(i interface{}) {
		w := i.(Worker)
		//fmt.Println("I am a worker")
		//fmt.Println(w)
		for i := 0; i < w.target_iterations; i++ {
			for {
				select {
				case <-w.worker_ctx.Done():
					w.worker_wg.Done()
				default:
					samples := w.run_iteration()
					// if !complete {
					// 	panic(fmt.Sprintf("Worker %d, unable to complete iteration %d", w.worker_id, i))
					// }\
					//fmt.Println(samples)
					w.samples_chan <- samples
				}
			}
		}
	}
}

// Run for Constant time
func RunForConstantTime() PoolFunc {
	return func(i interface{}) {
		w := i.(Worker)
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(w.target_duration))
		w.worker_ctx = ctx
		w.worker_cancel = cancel
		//fmt.Println("I am a worker")
		//fmt.Println(w)
		for {
			select {
			case <-w.worker_ctx.Done():
				w.worker_wg.Done()
			default:
				samples := w.run_iteration()
				// if !complete {
				// 	panic(fmt.Sprintf("Worker %d, unable to complete iteration %d", w.worker_id, i))
				// }	\
				w.samples_chan <- samples
			}
		}
	}
}

// Run without any context until killed
func RunNormally() PoolFunc {
	return func(i interface{}) {
		w := i.(Worker)
		for {
			select {
			case <-w.worker_ctx.Done():
				// kill worker
				w.worker_wg.Done()
			default:
				// perform each iteration
				samples := w.run_iteration()
				// push to global metrics channel
				w.samples_chan <- samples
			}
		}
	}
}

// RunOnce
func RunOnce() PoolFunc {
	return func(i interface{}) {
		w := i.(Worker)
		defer w.worker_wg.Done()
		samples := w.run_iteration()
		select {
		case w.samples_chan <- samples:
		default:
			// If the channel is full or unavailable, do nothing
			//fmt.Println("Samples channel is full or unavailable")
		}
	}
}
