package metrics

import (
	"context"
	"time"
)

// smallest data element

type Sample struct {
	metric   *Metric
	time     time.Time
	name     string
	val      float64
	metadata map[string]string
}

type Samples []Sample

type SampleContainer interface {
	GetSamples() Samples
}

var (
	_ SampleContainer = Samples{}
)

func (s Samples) GetSamples() Samples {
	return s
}

// fetched from a channel
func FetchBufferedSamples(input <-chan SampleContainer) []Samples {
	var res []Samples
	for {
		samples, ok := <-input
		if ok {
			s := samples.GetSamples()
			res = append(res, s)
		} else {
			return res
		}
	}
}

func PushToSampleContainer(ctx context.Context, input SampleContainer, sample_chan chan SampleContainer) bool {
	if ctx.Err() != nil {
		return false
	}
	sample_chan <- input
	return true
}
