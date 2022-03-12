package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	b := a[:1:2]

	fmt.Printf("a is %v\n", a)
	fmt.Printf("b is %v\n", b)

	b[0] = 10
	b = append(b, 3, 4)

	fmt.Printf("a is %v\n", a)
	fmt.Printf("b is %v\n", b)
}
