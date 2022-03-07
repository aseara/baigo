package main

import (
	"fmt"
	"github.com/aseara/baigo/lesson29/dumpintf"
	"unsafe"
)

func main() {
	var eif interface{} = dumpintf.T(5)
	var err error = dumpintf.T(5)

	fmt.Printf("T(5): %v\n", dumpintf.T(5))
	var s string
	fmt.Printf("s == \"\": %v\n", s == "")

	println("eif: ", eif)
	println("err: ", err)
	println("eif = err: ", eif == err)

	dumpintf.DumpEface(eif)
	dumpintf.DumpItabOfIface(unsafe.Pointer(&err))
	dumpintf.DumpDataOfIface(err)
}
