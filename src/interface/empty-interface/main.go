package main

import "fmt"

// 空接口是接口类型的特殊形式，空接口没有任何方法，因此任何类型都无须实现空接口。从实现的角度看，任何值都满足这个接口的需求。
// 因此空接口类型可以保存任何值，也可以从空接口中取出原值。

// 空接口的内部实现保存了对象的类型和指针。
// 使用空接口保存一个数据的过程会比直接用数据对应类型的变量保存稍慢。
// 因此在开发中，应在需要的地方使用空接口，而不是在所有地方使用空接口。

func main() {
	// 空接口的赋值如下：
	var any interface{}
	any = 1
	fmt.Println(any) // 1
	any = "hello"
	fmt.Println(any) // hello
	any = false
	fmt.Println(any) // false

	// 从空接口获取值
	// 保存到空接口的值，如果直接取出指定类型的值时，会发生编译错误
	// 声明a变量, 类型int, 初始值为1
	var a int = 1
	// 声明i变量, 类型为interface{}, 初始值为a, 此时i的值变为1
	var i interface{} = a
	// 声明b变量, 尝试赋值i
	// var b int = i // cannot use i (type interface {}) as type int in assignment: need type assertion
	var b int = i.(int) // 使用类型断言方式
	fmt.Println(b)      // 1

	// 空接口的值比较
	// 1) 类型不同的空接口间的比较结果不相同
	// 保存有类型不同的值的空接口进行比较时，Go 语言会优先比较值的类型。因此类型不同，比较结果也是不相同的
	// a保存整型
	var a1 interface{} = 100
	// b保存字符串
	var b1 interface{} = "100"
	// 两个空接口不相等
	fmt.Println(a1 == b1)
	// false
	// 2) 不能比较空接口中的动态值
	// 当接口中保存有动态类型的值时，运行时将触发错误，代码如下：
	// c保存包含10的整型切片
	var c interface{} = []int{10}
	// d保存包含20的整型切片
	var d interface{} = []int{20}
	// 这里会发生崩溃
	fmt.Println(c == d) // panic: runtime error: comparing uncomparable type []int
}
