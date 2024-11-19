package artillary

import "github.com/panjf2000/ants"

// This is the pool manager

type Stage struct {
	target_vus int
	duration   int
}

type Artillary struct {
	init_workers     int
	stages           []Stage
	cutoff_duration  int
	strategy_workers Strategy
	worker_pool 	 *ants.PoolWithFunc
}

func NewArtillary(start_workers int, stages []Stage, cutoff_duration int, strategy Strategy) *Artillary {
	var poolFunc PoolFunc 
	switch strategy {
	case shared_iterations:
		poolFunc = RunOnce()
		break
	case per_worker_iterations: 
		poolFunc = RunForIterations()
		break
	case constant_workers: 
		poolFunc = RunForConstantTime()
		break
	// case ramping_workers: 
	// 	break
	// case constant_arrival_rate: 
	// 	break 
	// case ramping_arrival_rate: 
	// 	break
	}
	pool, err := ants.NewPoolWithFunc(start_workers, poolFunc, ants.WithNonblocking(true), ants.WithPreAlloc(true))
	if err != nil {
		panic(err)
	}
	return &Artillary{
		init_workers: start_workers,
		stages: stages, 
		cutoff_duration: cutoff_duration,  
		strategy_workers: strategy,
		worker_pool: pool,
	}
}

func (a *Artillary) start_shared_iterations() {
	// add number of tasks to the pool
}

func (a *Artillary) start_per_worker_iterations() {
	// start the number of required workers 
}

func (a *Artillary) start_constant_workers() {
	// start the number of required workers which run for fixed duration
}

func StartArtillary() {
	

}
