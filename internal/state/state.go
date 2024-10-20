package internal

import "internal/metrics"

// should encapsulate the regs of the state and the current state of the tests

type TestState struct {
	// list of tests

	// channel for sending metrics
	metrics_pipe chan metrics.SampleContainer
	// list of thresholds for the corresponding tests
	//
}

type WorkerState struct {
	// init_workers
	// current_workers
}
