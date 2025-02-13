package metrics

import (
	"container/heap"
	// "fmt"
	"sort"
)

type MaxHeap []float64
type MinHeap []float64
type Datapoints []float64

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] } // max-heap
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any) {
	value, ok := x.(float64) 
	if !ok {
		panic("expected float64 value")
	}
	*h = append(*h, value)
}
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap) Len() int {return len(h)} 
func (h MinHeap) Less(i, j int) bool {return h[i] < h[j]} 
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(float64))
}
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianHeap struct {
	minHeap *MinHeap
	maxHeap *MaxHeap
}

func NewMedianHeap() *MedianHeap {
	return &MedianHeap{
		minHeap: &MinHeap{},
		maxHeap: &MaxHeap{},
	}
}

func (mh *MedianHeap) Add(value float64) {
	// Insert into maxHeap first
	heap.Push(mh.maxHeap, value)

	// Ensure balance: maxHeap can only contain the same or one more element than minHeap
	if mh.maxHeap.Len() > mh.minHeap.Len()+1 {
		heap.Push(mh.minHeap, heap.Pop(mh.maxHeap))
	}

	// Ensure all values in maxHeap are less than or equal to values in minHeap
	if mh.minHeap.Len() > 0 && (*mh.minHeap)[0] < (*mh.maxHeap)[0] {
		// Swap the tops of the heaps
		maxVal := heap.Pop(mh.maxHeap).(float64)
		minVal := heap.Pop(mh.minHeap).(float64)
		heap.Push(mh.maxHeap, minVal)
		heap.Push(mh.minHeap, maxVal)
	}

	// fmt.Printf("max heap size, %d\n", mh.maxHeap.Len()) 
	// fmt.Printf("min heap size, %d\n", mh.minHeap.Len())
}

func (mh *MedianHeap) FindMedian() float64 {
	if (mh.maxHeap.Len() == 0)  {
		return 0.0
	}
	if mh.maxHeap.Len() > mh.minHeap.Len() {
		return (*mh.maxHeap)[0]
	}
	median := ((*mh.minHeap)[0] + (*mh.maxHeap)[0]) / 2
	return median
}

func (d Datapoints) FindPercentile(p int16) float64 {
	sort.Float64s(d)
	n := len(d) 
	if n == 0 {
		return 0.0
	}
	idx := int(float64(p / 100.00) * float64(n - 1))
	return d[idx]
}