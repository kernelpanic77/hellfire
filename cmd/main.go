package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants"
	// "github.com/panjf2000/ants/v2"
)

// essentially for a load factor of 0.75, means that whenever the number of active workers are 0.75 * capacity, we would double to existing capacity

// now we need to make this thread safe because if multiple workers are calling the Fire and trying to resize it will cause a conflict

type army struct {
	ant_pool    *ants.PoolWithFunc
	work        chan int
	scaleFactor float32
	wg          *sync.WaitGroup
	lock        *sync.Mutex
}

func (a *army) Fire(args interface{}) error {
	// req := )
	a.lock.Lock()
	if a.ant_pool.Running() >= int(float32(a.scaleFactor)*float32(a.ant_pool.Cap())) {
		a.ant_pool.Tune(2 * a.ant_pool.Cap())
	}
	a.lock.Unlock()
	a.wg.Add(1)
	err := a.ant_pool.Invoke(args)
	return err
}

func main() {
	var wg sync.WaitGroup
	pool, err := ants.NewPoolWithFunc(100, func(i interface{}) {
		// fmt.Printf("%d Hello\n", i.(int))
		time.Sleep(20 * time.Second)
		// fmt.Printf("%d Bye\n", i.(int))
		wg.Done()
	}, ants.WithNonblocking(false))
	defer pool.Release()
	if err != nil {
		panic(err)
	}
	tasks := make(chan int, 1000)
	army := &army{ant_pool: pool, work: tasks, scaleFactor: 0.75, wg: &wg}

	start := time.Now()

	iterations := 1_000_000

	available = 200

	for i := 0; i < available; i++ {
		// fmt.Printf("Submitting %d, current capacity %d\n", i, army.ant_pool.Cap())
		err := army.Fire(i)
		if err != nil {
			panic(err)
		}
	}

	army.wg.Wait()
	end := time.Now()
	fmt.Printf("%f", end.Sub(start).Seconds())
}
