package main

import "fmt"

// new 只分配内存，而 make 只能用于 slice、map 和 channel 的初始化
func main() {
	// func new(Type) *Type
	// new 函数只接受一个参数，这个参数是一个类型，并且返回一个指向该类型内存地址的指针。同时 new 函数会把分配的内存置为零，也就是类型的零值。
	var sum *int
	sum = new(int)    //分配空间，返回的是指针
	fmt.Println(*sum) // 0
	*sum = 98
	fmt.Println(*sum) // 98
	// new 函数不仅仅能够为系统默认的数据类型，分配空间，自定义类型也可以使用 new 函数来分配空间:
	type Student struct {
		name string
		age  int
	}
	var s *Student
	s = new(Student) //分配空间，返回的是指针
	(*s).name = "dequan"
	fmt.Println(*s) // {dequan 0}

	// make 也是用于内存分配的，但是和 new 不同，它只用于 chan、map 以及 slice 的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，
	// 因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。

	// Go语言中的 new 和 make 主要区别如下：
	// make 只能用来分配及初始化类型为 slice、map、chan 的数据。new 可以分配任意类型的数据；
	// new 分配返回的是指针，即类型 *Type。make 返回引用，即 Type；
	// new 分配的空间被清零。make 分配空间后，会进行初始化；

}
