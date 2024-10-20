package artillary

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
}

func NewArtillary() {

}

func (a *Artillary) run_tests() {

}
