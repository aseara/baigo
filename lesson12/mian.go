package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var a, b = int(5), uint(6)
	var p uintptr = 0x12345678
	fmt.Println("signed integer a's length is", unsafe.Sizeof(a))   // 8
	fmt.Println("unsigned integer b's length is", unsafe.Sizeof(b)) // 8
	fmt.Println("uintptr's length is", unsafe.Sizeof(p))            // 8

	f := float32(139.8125)
	bits := math.Float32bits(f)
	fmt.Printf("%b\n", bits)

	s := `         ,_---~~~~~----._
_,,_,*^____      _____*g*\"*,--,
/ __/ /'     ^.  /      \ ^@q   f
[  @f | @))    |  | @))   l  0 _/
\/   \~____ / __ \_____/     \
|           _l__l_           I
}          [______]           I
]            | | |            |
]             ~ ~             |
|                            |
 |                           |`
	fmt.Println(s)
}
