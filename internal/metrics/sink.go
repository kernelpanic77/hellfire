package metrics

// Responsible for Aggregating Metrics
type Sink interface {
	AddSample(sample Sample)
}

var (
	_ Sink = &CounterSink{}
)

type CounterSink struct {
	Val float64
}

func (c *CounterSink) AddSample(s Sample) {
	c.Val += s.val
}

func NewSink(Type MetricType) Sink {
	var sink Sink
	switch metric_type := Type; metric_type {
	case Counter:
		sink = &CounterSink{Val: 0.0}
	}
	return sink
}
