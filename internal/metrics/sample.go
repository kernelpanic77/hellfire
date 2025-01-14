package metrics

import (
	"context"
	"time"
)

// smallest data element

type Sample struct {
	metric   []*Metric
	time     time.Time
	name     string
	val      float64
	// metadata map[string]string
}

func NewSample(metric []*Metric, time time.Time, name string, val float64) Sample {
	return Sample{
		metric: metric,
		time: time,  
		name: name, 
		val: val,
	}
}

type Samples []Sample


type SampleContainer interface {
	GetSamples() Samples
}

var (
	_ SampleContainer = Samples{}
	_ []SampleContainer = []SampleContainer{Samples{}}
)

func (s Samples) GetSamples() Samples {
	return s
}

// fetched from a channel
func FetchBufferedSamples(input <-chan SampleContainer) []Samples {
	var res []Samples
	curr_len := len(input)
	for i := 0; i < curr_len; i++ {
		samples, ok := <-input
		if(ok) {
			s := samples.GetSamples()
			res = append(res, s)
		}else {
			break
		}		
	}
	return res
}

func PushToSampleContainer(ctx context.Context, input SampleContainer, sample_chan chan SampleContainer) bool {
	if ctx.Err() != nil {
		return false
	}
	sample_chan <- input
	return true
}
