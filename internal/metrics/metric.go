package metrics

import (
	"context"
	"log"
	"time"
)

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
	Thresholds []*Threshold
	ctx context.Context
}

type MetricsMap map[string][]*Metric
type RegistryMap map[string]*MetricsRegistry

type MetricsRegistry struct {
	Metrics MetricsMap
}

var RegistryOfRegistry RegistryMap

func NewMetricsRegistry() *MetricsRegistry {
	metrics := make(MetricsMap)
	return &MetricsRegistry{
		Metrics: metrics,
	}
}

func (m *MetricsRegistry) FetchMetricByName(name string) []*Metric {
	metric, exists := m.Metrics[name] 
	if(!exists) {
		return nil	
	}
	return metric
} 

func (m *MetricsRegistry) RegisterMetric(metric_type MetricType, name string, val_type ValueType, tag string) {
	// var metric *Metric
	sink := NewSink(metric_type)
	metric := &Metric{Type: metric_type, Name: name, ValType: val_type, Sink: sink, Tag: tag}
	m.Metrics[name] = append(m.Metrics[name], metric)
}


func (m *Metric) checkThreshold(threshold *Threshold) {
	// check number of acceptable failure ratio, 0 by default 
	ratio := (threshold.failure_count / threshold.count_evals)
	if ratio >= int(threshold.failure_ratio) {
		// kill the Test
		log.Println(threshold.condition_text + " has failed!") 
		// kill the current test 
		
	}
}

func (m *Metric) observeMetric(threshold *Threshold) {
	ticker := time.NewTicker(threshold.observation_interval)
	for {
		select {
		case <-ticker.C: 
			curr_val := m.Sink.FetchSampleValue()
			validation := GetCondition[float64](threshold.aggregate_condition)
			valid := validation(curr_val, threshold.aggregate_val)
			threshold.count_evals++
			if !valid {
				threshold.failure_count += 1
			}			
		case <-m.ctx.Done(): 
			return 
		}
	}
}