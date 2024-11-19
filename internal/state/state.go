package internal

import "github.com/kernelpanic77/hellfire/internal/metrics"

// should encapsulate the regs of the state and the current state of the tests
// should strictly contain the results of the test and not the metadata
type TestState struct {
	// channel for sending metrics
	metrics_pipe chan metrics.SampleContainer
	// duration of the test so far 
	// number of scenarios completed 
	// number of iterations executed
	// curve of rise and fall of VUs so far 
	// Current Summary for the user
	// current state of artillary -> like the number of VUs which are running, available and pool size etc
}
