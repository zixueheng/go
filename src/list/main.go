package main

import (
	"container/list"
	"fmt"
)

// 列表使用 container/list 包来实现，内部的实现原理是双链表，列表能够高效地进行任意位置的元素插入和删除操作。

func main() {
	// 1) 通过 container/list 包的 New() 函数初始化 list
	// 变量名 := list.New()

	// 2) 通过 var 关键字声明初始化 list
	// var 变量名 list.List

	// 列表与切片和 map 不同的是，列表并没有具体元素类型的限制，因此，列表的元素可以是任意类型，这既带来了便利，也引来一些问题，
	// 例如给列表中放入了一个 interface{} 类型的值，取出值后，如果要将 interface{} 转换为其他类型将会发生宕机。

	l := list.New()
	l.PushBack("hello")
	l.PushFront(12)
	// 列表插入函数的返回值会提供一个 *list.Element 结构，这个结构记录着列表元素的值以及与其他节点之间的关系等信息，从列表中删除元素时，需要用到这个结构进行快速删除。
	el := l.PushBack("Element")
	// 在Element之后添加high
	l.InsertAfter("high", el)
	// 在Element之前添加noon
	l.InsertBefore("noon", el)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	// 12
	// hello
	// noon
	// Element
	// high

	l.Remove(el)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	// 12
	// hello
	// noon
	// high
}
