package main

import (
	"sync"

	"github.com/aseara/baigo/lesson27/trace"
)

// A1 method
func A1() {
	defer trace.Trace()()
	B1()
}

// B1 method
func B1() {
	defer trace.Trace()()
	C1()
}

// C1 method
func C1() {
	defer trace.Trace()()
	D()
	D()
}

// D method
func D() {
	defer trace.Trace()()
}

// A2 method
func A2() {
	defer trace.Trace()()
	B2()
}

// B2 method
func B2() {
	defer trace.Trace()()
	C2()
}

// C2 method
func C2() {
	defer trace.Trace()()
	D()
}

func main() {
	defer trace.Trace()()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		A2()
		wg.Done()
	}()

	A1()
	wg.Wait()
}
