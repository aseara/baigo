package main

import "github.com/aseara/baigo/lesson25"

// T type
type T struct{}

// M1 method
func (T) M1() {}

// M2 method
func (T) M2() {}

// M3 method
func (*T) M3() {}

// M4 method
func (*T) M4() {}

// S type
type S T

func main() {
	var n int
	lesson25.DumpMethodSet(n)
	lesson25.DumpMethodSet(&n)

	var t T
	lesson25.DumpMethodSet(t)
	lesson25.DumpMethodSet(&t)

	var s S
	lesson25.DumpMethodSet(s)
	lesson25.DumpMethodSet(&s)
}
