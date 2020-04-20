package main

import "fmt"

// 类型断言（Type Assertion）是一个使用在接口值上的操作，用于检查接口类型变量所持有的值是否实现了期望的接口或者具体的类型。
// 在Go语言中类型断言的语法格式如下：
// value, ok := x.(T)
// 其中，x 表示一个接口的类型，T 表示一个具体的类型（也可为接口类型）。

// 该断言表达式会返回 x 的值（也就是 value）和一个布尔值（也就是 ok），可根据该布尔值判断 x 是否为 T 类型：
//     如果 T 是具体某个类型，类型断言会检查 x 的动态类型是否等于具体类型 T。如果检查成功，类型断言返回的结果是 x 的动态值，其类型是 T。
//     如果 T 是接口类型，类型断言会检查 x 的动态类型是否满足 T。如果检查成功，x 的动态值不会被提取，返回值是一个类型为 T 的接口值。
//     无论 T 是什么类型，如果 x 是 nil 接口值，类型断言都会失败。

func main() {
	var x interface{}
	x = 10

	value, ok := x.(int)
	fmt.Println(value, ok)
	// 10 true
	
	// 用一个参数接收时
	v := x.(int)  //这里要确保 x 的动态类型是 他的实际类型 int，否则会panic
	fmt.Println(v) // 10

	// b := x.(string) // panic: interface conversion: interface {} is int, not string
	// fmt.Println(b)

	value1, ok1 := x.(float32)
	fmt.Println(value1, ok1)
	// 0 false

	// 类型断言还可以配合 switch 使用，这种格式只能在 switch 语句中使用，用在其他地方报错，如：
	// fmt.Println(x.(type)) // use of .(type) outside type switch
	switch x.(type) {
	case int:
		fmt.Println("the type of a is int")
	case string:
		fmt.Println("the type of a is string")
	case float64:
		fmt.Println("the type of a is float")
	default:
		fmt.Println("unknown type")
	}
	// the type of a is int
}
