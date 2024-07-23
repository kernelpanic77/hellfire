package hellfire

import (
	"fmt"
	"log"
	"time"

	"github.com/panjf2000/ants"
)

type Artillary struct {
	init_workers int
	active_pool  *ants.PoolWithFunc
	spawn_rate   int64
	target       int
	ticker       time.Ticker
	duration     int
}

type TaskMissile struct {
	idx     int
	task_fn task
	missile Missile
}

type task func(t *T)

type Missile func(t task)

func NewMissile(t task, dur int) Missile {
	return func(t task) {
		ticker := time.NewTicker(1 * time.Second)
		timer := time.NewTimer(time.Duration(dur) * time.Second)
		for range ticker.C {
			select {
			case <-timer.C:
				fmt.Println("Ho gaya!!!")
				ticker.Stop()
				break
			default:
				t(&T{})
			}
		}
	}
}

func FireMissile(i interface{}) {
	taskMissile := i.(TaskMissile)
	taskMissile.missile(taskMissile.task_fn)
}

func (a *Artillary) start(s *Scenario) {
	poolOptions := ants.WithNonblocking(true)
	pool, err := ants.NewPoolWithFunc(1000, FireMissile, poolOptions)
	if err != nil {
		fmt.Println("ajsldjas")
	}
	a.active_pool = pool
	a.init_workers = 1
	a.spawn_rate = s.Spawn_rate
	fmt.Println(s)
	a.target = int(s.Target)
	a.ticker = *time.NewTicker(time.Second * 1)
	a.duration = int(s.Duration)
	// fmt.Print("aijdlajld")
}

func (a *Artillary) Fire(task func(t *T), done chan bool) {
	cease_fire := make(chan bool)
	go func() {
		start := time.Now()
		for {
			select {
			case <-cease_fire:
				fmt.Println("Exited Main function!")
				done <- true
			case <-a.ticker.C:
				fmt.Println(a.active_pool.Running())
				elapsed := int(time.Since(start) / time.Second)
				if a.active_pool.Running() < a.target {
					// fmt.Println(a.spawn_rate)
					// add tasks equal to spawn rate
					for i := 0; i < int(a.spawn_rate) && a.active_pool.Running() < a.target; i++ {
						// var remaining int
						fmt.Println(elapsed)
						missile := NewMissile(task, a.duration-elapsed)
						taskMissile := TaskMissile{idx: a.active_pool.Running() + 1, missile: missile, task_fn: task}
						err := a.active_pool.Invoke(taskMissile)
						if err != nil {
							log.Println(err.Error())
						}
					}
				}
			}

		}
	}()
	time.Sleep(time.Second * time.Duration(a.duration))
	log.Println("Khatam tata bye bye!!!!")
	a.ticker.Stop()
	cease_fire <- true
}

func (a *Artillary) stop() {
	a.active_pool.Release()
}
