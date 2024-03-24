package main

import (
	"fmt"
	"sync"
	"time"
)

type signal struct{}

var _ready bool

func main() {
	fmt.Println("start a group of workers...")
	groupSignal := sync.NewCond(&sync.Mutex{})
	c := spawnGroup(worker, 5, groupSignal)

	time.Sleep(time.Second * 5)
	fmt.Println("the group of workers start to work...")

	groupSignal.L.Lock()
	_ready = true
	groupSignal.Broadcast()
	groupSignal.L.Unlock()

	<-c
	fmt.Println("the group of workers work done!")
}

func spawnGroup(f func(int), num int, groupSignal *sync.Cond) <-chan signal {
	c := make(chan signal)
	var wg sync.WaitGroup
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(i int) {
			groupSignal.L.Lock()
			for !_ready {
				groupSignal.Wait()
			}
			groupSignal.L.Unlock()
			fmt.Printf("worker %d: start to work...\n", i)
			f(i)
			wg.Done()
		}(i + 1)
	}

	go func() {
		wg.Wait()
		c <- signal(struct{}{})
	}()

	return c
}

func worker(i int) {
	fmt.Printf("worker %d: is working...\n", i)
	time.Sleep(time.Second * 1)
	fmt.Printf("worker %d: works done\n", i)
}
