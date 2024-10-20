package metrics

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	_ Endpoint = &Ingester{}
)

type SampleBuffer struct {
	mu     sync.Mutex
	buffer []SampleContainer
	maxLen int
}

func (sb *SampleBuffer) PushSamples(samples []SampleContainer) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	sb.buffer = append(sb.buffer, samples...)
}

func (sb *SampleBuffer) FetchSamples() []SampleContainer {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	samples, len_samples := sb.buffer, len(sb.buffer)
	if len_samples > sb.maxLen {
		sb.maxLen = len_samples
	}
	sb.buffer = make([]SampleContainer, 0, (sb.maxLen+len_samples)/2)
	return samples
}

type PeriodicFlusher struct {
	interval time.Time
	timer    time.Timer
	duration int
	callback func(interface{})
}

func (p *PeriodicFlusher) Flush(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stop Flushing Metrics")
			return
		case <-p.timer.C:
			p.callback()
		}
	}
}

type Ingester struct {
	// create a buffer to store the metrics
	buffer  *SampleBuffer
	flusher *PeriodicFlusher
}

func NewIngester() *Ingester {
	return &Ingester{
		buffer:  &SampleBuffer{},
		flusher: &PeriodicFlusher{},
	}
}

func (i *Ingester) Start(ctx context.Context) { // Should basically start a go routine which start a goroutine to flush metrics to respective Metrics Sink
	go i.flusher.Flush(ctx)
}

func (i *Ingester) AddSamples(samples []SampleContainer) { // Should push the samples to the buffer
	i.buffer.PushSamples(samples)
}

func (i *Ingester) send_to_sink() {
	// read from buffer
	sample_containers := i.buffer.FetchSamples()
	for _, container := range sample_containers {
		samples := container.GetSamples()
		for _, sample := range samples {
			// push to the sink of the sample
			m := sample.metric
			m.Sink.AddSample(sample)
		}
	}
}
