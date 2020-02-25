package main

import "fmt"

// 结构体可以包含一个或多个匿名（或内嵌）字段，即这些字段没有显式的名字，只有字段的类型是必须的，此时类型也就是字段的名字。
// 匿名字段本身可以是一个结构体类型，即结构体可以包含内嵌结构体。

type innerS struct {
	int1 int
	int2 int
}

type outerS struct {
	a      int
	b      float32
	int    // 匿名字段
	innerS // 匿名字段，结构体也是一种数据类型，所以它也可以作为一个匿名字段来使用
}

func main() {
	outer := new(outerS) //返回一个 *outerS 指针
	outer.a = 1
	outer.b = 1.1
	// 通过类型 outer.int 的名字来获取存储在匿名字段中的数据，于是可以得出一个结论：在一个结构体中对于每一种数据类型只能有一个匿名字段。
	outer.int = 2
	// 外层结构体通过 outer.int1 直接进入内层结构体的字段，内嵌结构体甚至可以来自其他包。
	outer.int1 = 3
	outer.int2 = 4

	fmt.Println(*outer)
	// {1 1.1 2 {3 4}}

	// 使用结构体字面量
	outer2 := outerS{10, 10.1, 20, innerS{30, 40}}
	fmt.Println(outer2)
	// {10 10.1 20 {30 40}}
	fmt.Printf("%+v", outer2) // 加号 标记（%+v）会添加字段名
	// {a:10 b:10.1 int:20 innerS:{int1:30 int2:40}}
}
