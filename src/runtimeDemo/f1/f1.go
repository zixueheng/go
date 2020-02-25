package f1

import (
	"fmt"
	"runtime"
)

// CallerExample0 runtime.Caller的示例
func CallerExample0() {
	pc, file, line, ok := runtime.Caller(0) // 0表示在当前找调用者
	if !ok {
		fmt.Println("Caller Failed")
	}
	funcName := runtime.FuncForPC(pc).Name() //调用者的函数名
	fmt.Println(funcName, file, line)
}

// CallerExample1 runtime.Caller的示例
func CallerExample1() {
	pc, file, line, ok := runtime.Caller(1) // 1 表示往上一层找调用者
	if !ok {
		fmt.Println("Caller Failed")
	}
	funcName := runtime.FuncForPC(pc).Name() //调用者的函数名
	fmt.Println(funcName, file, line)
}
