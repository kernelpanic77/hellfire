package artillary

import (
	"context"
	"internal/metrics"
)

// contains the general logic of creation of different workers
// ramp up policies etc

// shared iterations - eg: total 100 iterations shared between some VUs - simply create a pool of workers and add number of iterations to job pool
// (number of VUs to run concurrently)

// per vu iterations - each worker should run specific number of iterations

// constant VUs - for fixed duration

// ramping VUs up or down based on stages start with specific number of VUs

// constant-arrival-rate - increase the number of iterations to be executed every timeUNit for a fixed duration with a starting number of VUs and and maxVUs

// ramping-arrival-rate - ramping same thing with stages

// lets consider the iteration to be one execution by a worker

type Worker struct {
	worker_id         int
	target_iterations int                            // number of iterations worker must complete
	target_duration   int                            // target duration for the worker
	cutoff_duration   int                            // cutoff the worker after specific time
	iteration_func    func(interface{})              // function which worker is supposed to run
	worker_ctx        *context.Context               // context for each worker
	kill_worker       chan (bool)                    // kill worker channel
	worker_client     interface{}                    // should perform the heavy lifting of running an iteration and send metrics to the output channel
	samples_chan      chan []metrics.SampleContainer // dedicated channel for sending metrics
	strategy          Strategy                       // Strategy in workers
}

func NewWorker() *Worker {
	return nil
}

func (w *Worker) workerLoop() {
	// run until context doesn;t kill/cutoff time for fixed number of iterations/duration
	switch Strategy {
	case shared_iterations:
		err := w.run_shared_iterations()
	}
}

func (w *Worker) run_shared_iterations() error {
	return nil
}
