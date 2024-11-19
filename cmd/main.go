package main

import (
	"context"
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
	cancel_ctx	*[]context.CancelFunc
}

// graceful increase of capacity

// graceful decrease of capacity

func (a *army) kill_workers(count int) error {
	if(a.ant_pool.Running() >= count) {
		a.lock.Lock()
		for i := 0; i < count; i++ {
			if (*a.cancel_ctx)[i] != nil {
				fmt.Printf("Calling cancel function %d\n", i+1)
				(*a.cancel_ctx)[i]() // Invoke each cancel function
			}
		}
		a.lock.Unlock()
	}
	return nil
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

type worker struct {
	ctx context.Context
	id int
}

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	pool, err := ants.NewPoolWithFunc(100, func(i interface{}) {
		w := i.(worker) 
		for {
			select {
				case <- w.ctx.Done(): 
				fmt.Printf("Killing Worker %d", w.id)
				wg.Done()
				return
				default: 
				// fmt.Printf("Hello from %d", w.id)
				time.Sleep(time.Millisecond * 500)
				// fmt.Prin
			}
		}
	}, ants.WithNonblocking(false), ants.WithPreAlloc(true))
	defer pool.Release()
	if err != nil {
		panic(err)
	}
	fmt.Println(pool.Free())
	tasks := make(chan int, 1000)
	list_of_cancel_funcs := make([]context.CancelFunc, 0)
	army := &army{ant_pool: pool, work: tasks, scaleFactor: 0.75, wg: &wg, lock: &lock, cancel_ctx: &list_of_cancel_funcs}
	start := time.Now()
	available := 200000

	for i := 0; i < available; i++ {
		// fmt.Printf("Submitting %d, current capacity %d\n", i, army.ant_pool.Cap())
		ctx, cancel := context.WithCancel(context.Background())
		list_of_cancel_funcs = append(list_of_cancel_funcs, cancel)
		tmp_worker := worker{ctx: ctx, id: i}
		err := army.Fire(tmp_worker)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(pool.Running())
	army.kill_workers(5)
	fmt.Println(pool.Running())
	time.Sleep(1 * time.Second)
	fmt.Println(pool.Running())
	time.Sleep(1 * time.Second)
	fmt.Println(pool.Running())
	time.Sleep(1 * time.Second)
	fmt.Println(pool.Running())
	army.wg.Wait()
	end := time.Now()
	pool.Release()
	fmt.Printf("%f", end.Sub(start).Seconds())
}
