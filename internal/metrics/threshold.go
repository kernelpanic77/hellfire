package metrics

import (
	"log"
	"time"
)

type Threshold struct {
	// simple example could be expected latency of p95
	// can set a limit on count, value, rate
	// example p(99.9) < 200
	observation_interval time.Duration
	count_evals int 
	failure_count int
	failure_ratio  float32 
	delay_abort_eval int32
	condition_text      string
	aggregate           string
	aggregate_val       float64
	aggregate_condition string
	value               float64
}

type Number interface {
    int | float32 | float64
}

// Define a type for the comparison functions
type CompareFunc[T Number] func(a, b T) bool

func GetCondition[T Number](condition string) func(a, b T) bool {
	switch condition {
	case ">":
		return func(a, b T) bool { return a > b }
	case ">=":
		return func(a, b T) bool { return a >= b }
	case "<":
		return func(a, b T) bool { return a < b }
	case "<=":
		return func(a, b T) bool { return a <= b }
	case "==":
		return func(a, b T) bool { return a == b }
	case "!=":
		return func(a, b T) bool { return a != b }
	default:
		log.Fatalf("Invalid condition: %s", condition)
		return nil
	}
}
// parse a string to create a new threshold, store it with the metric

// create goroutines to monior metrics and thresholds and mark failures
// monitors the sink of each metric 



// count the number of times threshold is crossed












