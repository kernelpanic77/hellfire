package metrics

import "context"

// should encapsulate the entire metrics operations

// 1. channel which accepts metrics for all tests
// 2. constant monitoring of thresholds
// 3. creation of metrics
// 4. Preparation of Test summary

type Machine struct {
	// list of endpoints
	endpoints []Endpoint
	samples_chan chan SampleContainer
	registry     *MetricsRegistry
	sinks map[MetricType]Sink
	ctx context.Context
}

// initialize machine 
func NewMachine(ctx context.Context, registry *MetricsRegistry) *Machine {
	// create endpoints
	var endpoints []Endpoint 

	// create ingester
	ingester := NewIngester()
	endpoints = append(endpoints, ingester)

	// queue to ingest metrics 
	metrics_chan := make(chan SampleContainer, 1000000)

	// create sinks
	return &Machine{
		endpoints: endpoints,
		samples_chan: metrics_chan,
		registry: registry,
	}
}

func (m *Machine) GetEndpoints() []Endpoint{
	return m.endpoints
}

func (m *Machine) GetSamplesChan() chan SampleContainer {
	return m.samples_chan;
}