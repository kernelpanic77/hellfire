package artillary

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/kernelpanic77/hellfire/common"
	"github.com/kernelpanic77/hellfire/internal/metrics"
	"github.com/panjf2000/ants"
)

// This is the pool manager

type Stage struct {
	target_vus int
	duration   int
}

type Artillary struct {
	tag string
	init_workers     int
	stages           []Stage
	cutoff_duration  int
	strategy_workers Strategy
	SharedIterations int
	worker_pool 	 *ants.PoolWithFunc
	task common.Task
	ctx context.Context
	metrics_chan chan metrics.SampleContainer
	registry *metrics.MetricsRegistry
}

func NewArtillary(ctx context.Context, scenario common.Scenario, task common.Task, sampleschan chan metrics.SampleContainer) *Artillary {
	// find the max preallocated VUs needed for completing all the scenarios start that many initialy workers
	artillary := newArtillary(ctx, scenario, NewStrategy(scenario.Strategy), task, sampleschan)
	return artillary
}

func WithSharedIterations(artillary *Artillary, iterations int) *Artillary {
	artillary.SharedIterations = iterations
	return artillary
}

func newArtillary(ctx context.Context, scenario common.Scenario, strategy Strategy, task common.Task, metrics_chan chan metrics.SampleContainer) *Artillary {
	var poolFunc PoolFunc 
	total_iters := 0
	switch strategy {
	case shared_iterations:
		poolFunc = RunOnce()
		total_iters = int(scenario.Iterations)
	case per_worker_iterations: 
		poolFunc = RunForIterations()
	case constant_workers: 
		poolFunc = RunForConstantTime()
	// case ramping_workers: 
	// 	break
	// case constant_arrival_rate: 
	// 	break 
	// case ramping_arrival_rate: 
	// 	break
	}
	pool, err := ants.NewPoolWithFunc(int(scenario.PreAllocatedVUs), poolFunc, ants.WithNonblocking(true), ants.WithPreAlloc(true))
	if err != nil {
		panic(err)
	}
	return &Artillary{
		init_workers: int(scenario.PreAllocatedVUs),
		cutoff_duration: int(scenario.MaxDuration),  
		SharedIterations: total_iters,
		strategy_workers: strategy,
		worker_pool: pool,
		ctx: ctx,
		task: task,
		metrics_chan: metrics_chan,
	}
}

func (a *Artillary) start_shared_iterations(wg *sync.WaitGroup) {
	// add number of tasks to the pool
	wg.Add(1)
	go func(){
		defer wg.Done()
		for i := 0; i < a.SharedIterations; i++ {
			w := NewWorker(i, a.task, a.ctx, a.metrics_chan, wg)
			fmt.Println("triggering the test for worker")
			a.worker_pool.Invoke(w)
		}
	}()
}

func (a *Artillary) start_per_worker_iterations() {
	// start the number of required workers 
}

func (a *Artillary) start_constant_workers() {
	// start the number of required workers which run for fixed duration
}

func (a *Artillary) StartArtillary() {
	test_wg := &sync.WaitGroup{}
	test_wg.Add(1)
	timer := time.NewTimer(time.Duration(a.cutoff_duration * int(time.Second)))
	go func() {
		defer test_wg.Done()
		<-timer.C
		fmt.Println("Bas Khatam")
	}()
	switch a.strategy_workers {
	case shared_iterations: 
	a.start_shared_iterations(test_wg)
	}
	test_wg.Wait()
}
