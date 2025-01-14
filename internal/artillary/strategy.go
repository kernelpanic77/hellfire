package artillary

type Strategy int

const (
	shared_iterations Strategy = iota
	per_worker_iterations
	constant_workers
	ramping_workers
	constant_arrival_rate
	ramping_arrival_rate
)

func NewStrategy(strategy string) Strategy {
	var ret_strategy Strategy
	switch strategy {
	case "shared-iterations": 
	ret_strategy = shared_iterations
	case "per-worker-iteraitons": 
	ret_strategy = per_worker_iterations
	}
	return ret_strategy
}
