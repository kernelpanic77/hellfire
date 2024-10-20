package artillary

// This is the pool manager

type Stage struct {
	target_vus int
	duration   int
}

type Artillary struct {
	init_workers    int
	stages          []Stage
	cutoff_duration int
}
