package main

import (
	"fmt"
	"sync"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func echo(nums []int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, n := range nums {
			ch <- n
		}
		close(ch)
	}()
	return ch
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		sum := 0
		for n := range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

func merge(ins []<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(ins))
	for _, in := range ins {
		go func(in <-chan int) {
			for n := range in {
				out <- n
			}
			wg.Done()
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	nums := makeRange(1, 10000)
	in := echo(nums)

	const nProcess = 5
	var chans [nProcess]<-chan int
	for i := range chans {
		chans[i] = sum(in)
	}

	for n := range sum(merge(chans[:])) {
		fmt.Printf("sum is %05d\n", n)
	}

}
