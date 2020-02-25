package main

import "fmt"

// Go 方法是作用在接收器（receiver）上的一个函数，接收器是某种类型的变量，因此方法是一种特殊类型的函数。

// Bag 使用背包作为“对象”，将物品放入背包的过程作为“方法”
type Bag struct {
	items []int // 整型切片类型的 items 的成员
}

// 1、为结构体添加方法

// Insert 将一个物品放入背包的过程
func Insert(b *Bag, itemid int) {
	b.items = append(b.items, itemid)
}

// 2、Go语言的结构体方法

// Insert2 将背包及放入背包的物品中使用Go语言的结构体和方法方式编写，为 *Bag 创建一个方法
// Insert2(itemid int) 的写法与函数一致，(b*Bag) 表示接收器，即 Insert2 作用的对象实例
func (b *Bag) Insert2(itemid int) {
	b.items = append(b.items, itemid)
}

// 接收器的格式如下：
// func (接收器变量 接收器类型) 方法名(参数列表) (返回参数) {
//     函数体
// }

// 1、指针类型的接收器由一个结构体的指针组成，更接近于面向对象中的 this 或者 self。
// 由于指针的特性，调用方法时，修改接收器指针的任意成员变量，在方法结束后，修改都是有效的。

// 2、当方法作用于非指针接收器时，Go语言会在代码运行时将接收器的值复制一份，在非指针接收器的方法中可以获取接收器的成员值，但修改后无效。

// 3、指针和非指针接收器的使用
// 在计算机中，小对象由于值复制时的速度较快，所以适合使用非指针接收器，大对象因为复制性能较低，适合使用指针接收器，在接收器和参数间传递时不进行复制，只是传递指针。

func main() {
	bag := new(Bag)  // 返回Bag类型的指针
	Insert(bag, 100) // 传统面向过程的调用方式
	fmt.Println((*bag).items)
	// [100]

	// 在 Insert() 转换为方法后，我们就可以愉快地像其他语言一样，用面向对象的方法来调用 b 的 Insert。
	bag.Insert2(101)
	fmt.Println((*bag).items)
	// [100 101]
}
