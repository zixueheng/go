package main

import "fmt"

// Go语言的结构体内嵌特性就是一种组合特性，使用组合特性可以快速构建对象的不同特性

// Flying 可飞行的
type Flying struct{}

// Fly 结构体Flying方法
func (f *Flying) Fly() {
	fmt.Println("can fly")
}

// Walkable 可行走的
type Walkable struct{}

// Walk 结构体Walkable方法
func (f *Walkable) Walk() {
	fmt.Println("can calk")
}

// Human 人类，嵌入可行走结构（Walkable），让人类具备“可行走”特性
type Human struct {
	Walkable // 人类能行走
}

// Bird 鸟类，嵌入可行走结构（Walkable）和可飞行结构（Flying），让鸟类具备既可行走又可飞行的特性。
type Bird struct {
	Walkable // 鸟类能行走
	Flying   // 鸟类能飞行
}

func main() {
	// 实例化鸟类
	b := new(Bird)
	fmt.Println("Bird: ")
	// 调用鸟类可以使用的功能，如飞行和行走
	b.Fly()
	b.Walk()
	// Bird:
	// can fly
	// can calk

	// 实例化人类
	h := new(Human)
	fmt.Println("Human: ")
	// 调用人类能使用的功能，如行走
	h.Walk()
	// Human:
	// can calk
}
