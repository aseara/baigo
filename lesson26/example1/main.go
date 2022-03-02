package main

import "fmt"

// Interface1 intf
type Interface1 interface {
	M1()
}

// Interface2 intf
type Interface2 interface {
	M1()
	M2()
}

// Interface3 intf
type Interface3 interface {
	Interface1
	Interface2 // Error: duplicate method M1
}

// Interface4 intf
type Interface4 interface {
	Interface2
	M2() // Error: duplicate method M2
}

func main() {
	var intf3 Interface3 = struct {
		Interface3
	}{}
	fmt.Printf("type of intf3: %T\n", intf3)

	var intf4 Interface3 = struct {
		Interface4
	}{}
	fmt.Printf("type of intf4: %T\n", intf4)

	type S struct {
		Interface3
	}
	var s1 S
	fmt.Printf("type of s1: %T\n", s1)
}
