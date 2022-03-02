package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

func main() {
	var s = "中国人"
	fmt.Printf("the length of s = %d\n", len(s)) // 9

	for i := 0; i < len(s); i++ {
		fmt.Printf("0x%x ", s[i]) // 0xe4 0xb8 0xad 0xe5 0x9b 0xbd 0xe4 0xba 0xba
	}
	fmt.Printf("\n")

	fmt.Println("the character count in s is", utf8.RuneCountInString(s)) // 3
	for _, c := range s {
		fmt.Printf("0x%x ", c)
	}
	fmt.Printf("\n")

	s = "hello"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("Data: 0x%x, Len: %d\n", hdr.Data, hdr.Len)
	p := (*[5]byte)(unsafe.Pointer(hdr.Data))
	dumpBytes((*p)[:])
}

func dumpBytes(arr []byte) {
	fmt.Printf("[")
	for _, b := range arr {
		fmt.Printf("%c ", b)
	}
	fmt.Printf("]\n")
}
