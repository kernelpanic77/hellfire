package hellfire

import "testing"

type Stage struct {
	Target   int64
	Duration int64
}

type Scenario struct {
	Name            string
	PreAllocatedVUs int64
	Spawn_rate      int64
	maxVUs          int64
	Stages          []Stage
	maxDuration     int64
}

type Threshold struct {
	ScenarioName string
	MetricName   string
	Conditions   []string
}

type TestMetadata struct {
	Scenarios  []Scenario
	Thresholds []Threshold
	T          *testing.T
	Iteration  task
}
