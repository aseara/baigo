package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/aseara/baigo/lesson27/trace/instrumenter"
	"github.com/aseara/baigo/lesson27/trace/instrumenter/ast"
)

var _wrote bool

func init() {
	flag.BoolVar(&_wrote, "w", false, "write result to (source) file instead of output")
}

func usage() {
	fmt.Println("instrument [-w] xxx.go")
	flag.PrintDefaults()
}

func main() {
	fmt.Println(os.Args)
	flag.Usage = usage
	flag.Parse()

	if len(os.Args) < 2 {
		usage()
		return
	}

	var file string
	if len(os.Args) == 3 {
		file = os.Args[2]
	}

	if len(os.Args) == 2 {
		file = os.Args[1]
	}
	if filepath.Ext(file) != ".go" {
		usage()
		return
	}

	var ins instrumenter.Instrumenter

	ins = ast.New("github.com/aseara/baigo/lesson27/trace", "trace", "Trace")
	newSrc, err := ins.Instrument(file)
	if err != nil {
		panic(err)
	}

	if newSrc == nil {
		fmt.Printf("no trace added for %s\n", file)
	}

	if !_wrote {
		fmt.Println(string(newSrc))
		return
	}

	if err = ioutil.WriteFile(file, newSrc, 0666); err != nil {
		fmt.Printf("write %s error: %v\n", file, err)
		return
	}

	fmt.Printf("instrument trace for %s ok\n", file)
}
