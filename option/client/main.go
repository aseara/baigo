package main

import "github.com/aseara/baigo/option/demo"

func main() {
	var foo demo.Foo
	preVerbosity := foo.Option(demo.Verbosity(3))
	foo.Option(preVerbosity)
}
