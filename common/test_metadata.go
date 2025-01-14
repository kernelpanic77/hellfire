package common

import "testing"

// Stage struct
type Stage struct {
	Target   int64
	Duration int64
}

func (s *Stage) GetTarget() int64 {
	return s.Target
}

func (s *Stage) SetTarget(target int64) {
	s.Target = target
}

func (s *Stage) GetDuration() int64 {
	return s.Duration
}

func (s *Stage) SetDuration(duration int64) {
	s.Duration = duration
}

// Scenario struct
type Scenario struct {
	Strategy string
	Iterations int64
	Name            string
	PreAllocatedVUs int64
	Spawn_rate      int64
	MaxVUs          int64
	Stages          []Stage
	MaxDuration     int64
}

func (sc *Scenario) GetName() string {
	return sc.Name
}

func (sc *Scenario) SetName(name string) {
	sc.Name = name
}

func (sc *Scenario) GetPreAllocatedVUs() int64 {
	return sc.PreAllocatedVUs
}

func (sc *Scenario) SetPreAllocatedVUs(preAllocatedVUs int64) {
	sc.PreAllocatedVUs = preAllocatedVUs
}

func (sc *Scenario) GetSpawnRate() int64 {
	return sc.Spawn_rate
}

func (sc *Scenario) SetSpawnRate(spawnRate int64) {
	sc.Spawn_rate = spawnRate
}

func (sc *Scenario) GetMaxVUs() int64 {
	return sc.MaxVUs
}

func (sc *Scenario) SetMaxVUs(maxVUs int64) {
	sc.MaxVUs = maxVUs
}

func (sc *Scenario) GetStages() []Stage {
	return sc.Stages
}

func (sc *Scenario) SetStages(stages []Stage) {
	sc.Stages = stages
}

func (sc *Scenario) GetMaxDuration() int64 {
	return sc.MaxDuration
}

func (sc *Scenario) SetMaxDuration(maxDuration int64) {
	sc.MaxDuration = maxDuration
}

// Threshold struct
type Threshold struct {
	ScenarioName string
	MetricName   string
	Conditions   []string
}

func (th *Threshold) GetScenarioName() string {
	return th.ScenarioName
}

func (th *Threshold) SetScenarioName(scenarioName string) {
	th.ScenarioName = scenarioName
}

func (th *Threshold) GetMetricName() string {
	return th.MetricName
}

func (th *Threshold) SetMetricName(metricName string) {
	th.MetricName = metricName
}

func (th *Threshold) GetConditions() []string {
	return th.Conditions
}

func (th *Threshold) SetConditions(conditions []string) {
	th.Conditions = conditions
}

// TestMetadata struct
type TestMetadata struct {
	Scenarios  []Scenario
	Thresholds []Threshold
	T          *testing.T
	Iteration  Task
}

func (tm *TestMetadata) GetScenarios() []Scenario {
	return tm.Scenarios
}

func (tm *TestMetadata) SetScenarios(scenarios []Scenario) {
	tm.Scenarios = scenarios
}

func (tm *TestMetadata) GetThresholds() []Threshold {
	return tm.Thresholds
}

func (tm *TestMetadata) SetThresholds(thresholds []Threshold) {
	tm.Thresholds = thresholds
}

func (tm *TestMetadata) GetT() *testing.T {
	return tm.T
}

func (tm *TestMetadata) SetT(t *testing.T) {
	tm.T = t
}

func (tm *TestMetadata) GetIteration() Task {
	return tm.Iteration
}

func (tm *TestMetadata) SetIteration(iteration Task) {
	tm.Iteration = iteration
}
