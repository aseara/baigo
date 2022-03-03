package dumpintf

import (
	"fmt"
	"unsafe"
)

const ptrSize = unsafe.Sizeof(uintptr(0))

// DumpEface 输出 interface{} 变量的内部信息
func DumpEface(i interface{}) {
	ptrToEface := (*eface)(unsafe.Pointer(&i))
	fmt.Printf("eface: %+v\n", *ptrToEface)
	if ptrToEface._type != nil {
		// dump _type info
		fmt.Printf("     -type: %+v\n", *(ptrToEface._type))
	}
	if ptrToEface.data != nil {
		switch i.(type) {
		case int:
			dumpInt(ptrToEface.data)
		case float64:
			dumpFloat64(ptrToEface.data)
		case T:
			dumpT(ptrToEface.data)
		default:
			fmt.Printf("     unsupported data type\n")
		}
	}
	fmt.Printf("\n")
}

// DumpItabOfIface 输出 iface 中 itab 的详细信息
func DumpItabOfIface(ptrToIface unsafe.Pointer) {
	p := (*iface)(ptrToIface)
	fmt.Printf("iface: %+v\n", *p)
	if p.tab != nil {
		// dump itab
		fmt.Printf("     itab: %+v\n", *(p.tab))
		// dump inter in itab
		fmt.Printf("        inter: %+v\n", *(p.tab.inter))
		// dump _type in itab
		fmt.Printf("        _type: %+v\n", *(p.tab._type))
		// dump fun in itab
		funPtr := unsafe.Pointer(&(p.tab.fun))
		fmt.Printf("        fun: [")
		for i := 0; i < len((*(p.tab.inter)).mhdr); i++ {
			tp := (*uintptr)(unsafe.Pointer(uintptr(funPtr) + uintptr(i)*ptrSize))
			fmt.Printf("0x%x(%d)", *tp, *tp)
		}
		fmt.Printf("]\n")
	}

}

// DumpDataOfIface 打印 iface data部分的详细信息
func DumpDataOfIface(i interface{}) {
	// this is a trick as the data part of eface and iface are same
	ptrToEface := (*eface)(unsafe.Pointer(&i))
	if ptrToEface.data != nil {
		// dump data
		switch i.(type) {
		case int:
			dumpInt(ptrToEface.data)
		case float64:
			dumpFloat64(ptrToEface.data)
		case T:
			dumpT(ptrToEface.data)
		default:
			fmt.Printf("     unsupported data type\n")
		}
	}
	fmt.Printf("\n")
}

func dumpT(data unsafe.Pointer) {
	p := (*T)(data)
	fmt.Printf("     data: %+v", *p)
}

func dumpFloat64(data unsafe.Pointer) {
	p := (*float64)(data)
	fmt.Printf("     data: %f\n", *p)
}

func dumpInt(data unsafe.Pointer) {
	p := (*int)(data)
	fmt.Printf("     data: %d\n", *p)
}
