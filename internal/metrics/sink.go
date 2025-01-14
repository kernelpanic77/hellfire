package metrics

import "math"

// Responsible for Aggregating Metrics
type Sink interface {
	AddSample(sample Sample)
	FetchSampleValue() float64
}

var (
	_ Sink = &CounterSink{}
	_ Sink = &GuageSink{}
	_ Sink = &TrendSink{} 
	_ Sink = &RateSink{}
)
type CounterSink struct {
	Val float64
}

func (c *CounterSink) AddSample(s Sample) {
	c.Val += s.val
}

func (c *CounterSink) FetchSampleValue() float64 {
	return c.Val
}

type GuageSink struct {
	MinVal float64 
	MaxVal float64 
	RecentVal float64
}

func (g *GuageSink) AddSample(s Sample) {
	g.MinVal = math.Min(g.MinVal, s.val) 
	g.MaxVal = math.Max(g.MaxVal, s.val) 
	g.RecentVal = s.val
}

func (g *GuageSink) FetchSampleValue() float64 {
	return 0.0
}

func (g *GuageSink) FetchGuage() []float64 {
	return []float64{g.MinVal, g.MaxVal}
}

type TrendSink struct {
	datapoints Datapoints
	medianHeap MedianHeap
	cumulativeSum float64
	MeanVal float64 
	MedianVal float64 
	P90Val float64
	P95Val float64 
}

func (t *TrendSink) AddSample(s Sample) {
	t.cumulativeSum += s.val 
	t.datapoints = append(t.datapoints, s.val) 
	t.MeanVal = t.cumulativeSum / float64(len(t.datapoints)) 
	t.medianHeap.Add(s.val) 
}

func (t *TrendSink) FetchSampleValue() float64 {
	return 0.0
}

func (t *TrendSink) FetchTrends() []float64 {
	return []float64{t.MeanVal, t.medianHeap.FindMedian(), t.datapoints.FindPercentile(90), t.datapoints.FindPercentile(95)}
}
type RateSink struct {
	Val float64 
	count int64 
	sum int64 
}

func (r *RateSink) AddSample(s Sample) {
	// should this be atomic
	r.count += 1 
	r.sum += int64(s.val)
	r.Val = float64(r.sum / r.count) * 100.00 
}

func (r *RateSink) FetchSampleValue() float64 {
	return r.Val 
}

func NewSink(Type MetricType) Sink {
	var sink Sink
	switch metric_type := Type; metric_type {
	case Counter:
		sink = &CounterSink{Val: 0.0}
	case Guage: 
		sink = &GuageSink{MaxVal: math.MinInt64, MinVal: math.MaxInt64, RecentVal: math.MinInt64}
	case Trend: 
		sink = &TrendSink{datapoints: make([]float64, 0), MeanVal: 0.0, MedianVal: 0.0, P90Val: 0.0, P95Val: 0.0}
	case Rate: 
		sink = &RateSink{Val: 0.0, count: 0, sum: 0}
	}
	return sink
}
