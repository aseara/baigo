package main

import "fmt"

// 测试生成 WebAssembly 文件
// $env:GOARCH="wasm"
// $env:GOOS="js"
// go build -o lib.wasm main.go
func main() {
	fmt.Println("hello, webAssembly")
}
