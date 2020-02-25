package main

import (
	"fmt"
	"runtime"

	"./f1"
)

func main() {
	_, file, line, ok := runtime.Caller(0) // file表示调用的文件，line表示调用的行数
	if !ok {
		fmt.Println("Caller Failed")
	}
	fmt.Println(file, line) // D:/Go/src/runtime/main.go 9

	f1.CallerExample0() // _/D_/Go/src/runtimeDemo/f1.CallerExample0 D:/Go/src/runtimeDemo/f1/f1.go 10
	f1.CallerExample1() // main.main D:/Go/src/runtimeDemo/main.go 20
}
