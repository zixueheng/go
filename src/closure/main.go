package main

import "fmt"

// 闭包是一个能包含外部作用域的变量的函数
// func adder(x int) func(y int) int {
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }

// func main() {
// 	fmt.Println(adder(100)(200))
// }

// 示例2
func f1(f func()) {
	fmt.Println("F1")
	f()
}

func f2(x, y int) {
	fmt.Println("F2")
	fmt.Println(x + y)
}

// 现在要实现 f1 中调用 f2 即 f1(f2)，封装一个 f3 返回值类型和 f1 参数的一致（即 f3能作为f1的参数），在 f3 中执行 f2
func f3(f func(int, int), x, y int) func() { //这里的参数 f 类型和 函数 f2 一致
	return func() {
		f(x, y) //直接调用 f 函数 即调用 f2
	}
}

func main() {
	f1(f3(f2, 100, 200))
	// F1
	// F2
	// 300
}
