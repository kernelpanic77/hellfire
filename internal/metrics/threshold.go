package metrics

type Threshold struct {
	// simple example could be expected latency of p95
	// can set a limit on count, value, rate
	// example p(99.9) < 200

	condition_text      string
	aggregate           string
	aggregate_val       float32
	aggregate_condition string
	value               float32
}

// parse a string to create a new threshold, store it with the metric

// create goroutines to monior metrics and thresholds and mark failures
