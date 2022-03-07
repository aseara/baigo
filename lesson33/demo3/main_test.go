package main

import (
	"sync"
	"testing"
)

type counter1 struct {
	mu sync.Mutex
	i  int
}

var _cter counter1

func increase1() int {
	_cter.mu.Lock()
	defer _cter.mu.Unlock()
	_cter.i++
	return _cter.i
}

func BenchmarkIncrease1(b *testing.B) {
	var (
		wg  sync.WaitGroup
		num = 1000_000
	)

	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			_ = increase1()
			wg.Done()
		}(i)
	}
	wg.Wait()
}

type counter2 struct {
	c chan int
	i int
}

func newCounter2() *counter2 {
	cter := &counter2{
		c: make(chan int),
	}
	go func() {
		for {
			cter.i++
			cter.c <- cter.i
		}
	}()
	return cter
}

func (cter *counter2) increase2() int {
	return <-cter.c
}

func BenchmarkIncrease2(b *testing.B) {
	var (
		wg   sync.WaitGroup
		cter = newCounter2()
		num  = 1000_000
	)

	wg.Add(num)
	for i := 0; i < num; i++ {
		go func(i int) {
			_ = cter.increase2()
			wg.Done()
		}(i)
	}
	wg.Wait()
}
