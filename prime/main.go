package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("args error")
		return
	}

	p, err := strconv.Atoi(os.Args[1])
	if err != nil || p < 2 {
		fmt.Printf("args : %s", os.Args[1])
	}

	fmt.Printf("%d is prime: %v\n", p, prime(p))

	cache := []int{2, 3}
	for _, n := range cache {
		fmt.Printf(" %d", n)
	}
	fmt.Println()
}

func prime(c int) bool {
	if c == 2 || c == 3 {
		return true
	}

	cache := []int{2, 3}
	e := int(math.Sqrt(float64(c)))

	fmt.Printf("end of check: %d\n", e)

outer:
	for i := 5; i <= e; i += 2 {
		for _, p := range cache {
			if i%p == 0 {
				continue outer
			}
		}
		cache = append(cache, i)
	}

	for _, p := range cache {
		if c%p == 0 {
			fmt.Printf("facotr: %d\n", p)
			return false
		}
	}

	return true
}
