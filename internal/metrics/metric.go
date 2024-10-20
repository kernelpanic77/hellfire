package metrics

type MetricType int

type ValueType int // can be either time or data

const (
	Counter MetricType = iota
	Guage
	Rate
	Trend
)

const (
	Time ValueType = iota
	Data
)

type Metric struct {
	Type    MetricType
	Name    string
	ValType ValueType
	Sink    Sink
	Tag     string
}

type MetricsRegistry struct {
	Metrics map[string]*Metric
}

func (m *MetricsRegistry) NewMetrics(metric_type MetricType, name string, val_type ValueType, sink Sink, tag string) *Metric {
	_, ok := m.Metrics[name]
	if !ok {
		// var metric *Metric
		metric := &Metric{Type: metric_type, Name: name, ValType: val_type, Sink: sink, Tag: tag}
		m.Metrics[name] = metric
		return metric
	}
	return nil
}
