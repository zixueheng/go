package main

import "fmt"

// Go语言的类型或结构体没有构造函数的功能，但是我们可以使用结构体初始化的过程来模拟实现构造函数。

// 1、多种方式创建和初始化结构体——模拟构造函数重载

// Cat 结构体描述猫的特性
type Cat struct {
	Color string
	Name  string
}

// NewCatByName 定义用名字构造猫结构的函数，返回 Cat 指针。
func NewCatByName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// NewCatByColor 定义用颜色构造猫结构的函数，返回 Cat 指针。
func NewCatByColor(color string) *Cat {
	return &Cat{
		Color: color,
	}
}

func main() {
	cat1 := NewBlackCat("black")
	fmt.Println((*cat1))
}

// 2、带有父子关系的结构体的构造和初始化——模拟父级构造调用
// Cat 结构体类似于面向对象中的“基类”，BlackCat 嵌入 Cat 结构体，类似于面向对象中的“派生”，实例化时，BlackCat 中的 Cat 也会一并被实例化。

// BlackCat 嵌入Cat, 类似于派生，拥有 Cat 的所有成员，实例化后可以自由访问 Cat 的所有成员。
type BlackCat struct {
	Cat
}

// NewCat “构造基类”
func NewCat(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

// NewBlackCat “构造子类”
func NewBlackCat(color string) *BlackCat {
	cat := &BlackCat{} // 实例化 BlackCat 结构，此时 Cat 也同时被实例化
	cat.Color = color  // 填充 BlackCat 中嵌入的 Cat 颜色属性，BlackCat 没有任何成员，所有的成员都来自于 Cat
	return cat
}
