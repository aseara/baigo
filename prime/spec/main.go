package main

import "fmt"

func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(src <-chan int, dst chan<- int, prime int) {
	for i := range src {
		if i%prime != 0 {
			dst <- i
		}
	}
}

func sieve() {
	ch := make(chan int)
	go generate(ch)
	for {
		prime := <-ch
		fmt.Printf("  %5d\n", prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}

func main() {
	sieve()
}
