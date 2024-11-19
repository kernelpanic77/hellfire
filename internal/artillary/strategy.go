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


