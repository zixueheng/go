package main

import "fmt"

// Go语言可以将类型的方法与普通函数视为一个概念，从而简化方法和函数混合作为回调类型时的复杂性。
// 这个特性和 C# 中的代理（delegate）类似，调用者无须关心谁来支持调用，系统会自动处理是否调用普通函数或类型的方法。

// 方法和函数的统一调用

// 声明一个结构体
type class struct {
}

// 给结构体添加Do方法
func (c *class) Do(v int) {
	fmt.Println("call method do:", v)
}

// 普通函数的Do
func funcDo(v int) {
	fmt.Println("call function do:", v)
}
func main() {
	// 声明一个函数回调
	var delegate func(int)
	// 创建结构体实例
	c := new(class)
	// 将回调设为c的Do方法
	delegate = c.Do
	// 调用
	delegate(100)
	// call method do: 100

	// 将回调设为普通函数
	delegate = funcDo
	// 调用
	delegate(101)
	// call function do: 101

	// 这段代码能运行的基础在于：无论是普通函数还是结构体的方法，只要它们的签名一致，与它们签名一致的函数变量就可以保存普通函数或是结构体方法。
}
