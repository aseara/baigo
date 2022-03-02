package trace

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"
)

var _goroutineSapce = []byte("goroutine ")
var _mu sync.Mutex
var _m = make(map[uint64]int)

// Trace for tracing using defer
func Trace() func() {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		panic("not found caller")
	}

	fn := runtime.FuncForPC(pc)
	name := fn.Name()

	gid := curGoroutineID()

	_mu.Lock()
	indent := _m[gid] + 1
	_m[gid] = indent
	_mu.Unlock()

	printTrace(gid, name, "->", indent)
	return func() {
		_mu.Lock()
		_m[gid] = indent - 1
		_mu.Unlock()
		printTrace(gid, name, "<-", indent)
	}
}

func curGoroutineID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, _goroutineSapce)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return n
}

func printTrace(id uint64, name, arrow string, indent int) {
	indents := ""
	for i := 0; i < indent; i++ {
		indents += "    "
	}
	fmt.Printf("g[%05d]:%s%s%s\n", id, indents, arrow, name)
}
