package metrics

// should encapsulate the entire metrics operations

// 1. channel which accpet metrics for all tests
// 2. constant monitoring of thresholds
// 3. creation of metrics
// 4. Preparation of Test summary

type Machine struct {
	ingester     *Ingester
	samples_chan chan []SampleContainer
	registry     *MetricsRegistry
}
