package main

import (
	"fmt"
	"time"
)

// func main() {
// 	defer greet()()
// 	fmt.Println("Some code here...")
// }

// // 父函数返回的函数将是实际的延迟函数。父函数中的其他代码将在函数开始时（由 defer 语句放置的位置决定）立即执行。
// func greet() func() {
// 	fmt.Println("Hello!")
// 	// 返回的函数将作为真正的延迟函数
// 	return func() {
// 		fmt.Println("Bye!")
// 	} // this will be deferred
// }

// Hello!
// Some code here...
// Bye!

func main() {
	example()
	otherExample()
}

func example() {
	defer measure("example")()
	fmt.Println("Some code here")
}

func otherExample() {
	defer measure("otherExample")()
	fmt.Println("Some other code here")
}

// 在函数内定义的匿名函数可以访问完整的词法环境（lexical environment），这意味着在函数中定义的内部函数可以引用该函数的变量
func measure(name string) func() {
	start := time.Now() // start 参数变量在 measure 函数第一次执行和其延迟执行的子函数内都能访问到
	fmt.Printf("Starting function %s on %s\n", name, start)
	return func() {
		fmt.Printf("Exiting function %s after %s\n", name, time.Since(start))
	}
}

// Starting function example on 2019-12-17 10:47:40.3709361 +0800 CST m=+0.009965101
// Some code here
// Exiting function example after 158.5772ms
// Starting function otherExample on 2019-12-17 10:47:40.5295133 +0800 CST m=+0.168542301
// Some other code here
// Exiting function otherExample after 1.0044ms
