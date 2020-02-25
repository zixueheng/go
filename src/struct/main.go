package main

import "fmt"

type person struct {
	name string
	age  int
}

func f1(p person) {
	p.age = 18
}
func f2(p *person) {
	(*p).age = 20
}

// Command 使用结构体定义一个命令行指令，指令中包含名称、变量关联和注释等
type Command struct {
	Name    string // 指令名称
	Var     *int   // 指令绑定的变量
	Comment string // 指令的注释
}

func newCommand(name string, varref *int, comment string) *Command {
	return &Command{
		Name:    name,
		Var:     varref,
		Comment: comment,
	}
}

func main() {
	// 结构体本身是一种类型，可以像整型、字符串等类型一样，以 var 的方式声明结构体即可完成实例化
	var p1 person
	p1.name = "xiaoming"
	p1.age = 17
	fmt.Println(p1)
	f1(p1) //Go函数调用 的参数是拷贝
	fmt.Println(p1)
	f2(&p1)
	fmt.Println(p1)
	// {xiaoming 17}
	// {xiaoming 17}
	// {xiaoming 20}

	// new 关键字对类型（包括结构体、整型、浮点数、字符串等）进行实例化，结构体在实例化后会形成指针类型的结构体。
	// new 只分配内存，而 make 只能用于 slice、map 和 channel 的初始化
	p2 := new(person) //返回 person类型的指针
	fmt.Printf("%T, %v\n", p2, p2)
	(*p2).name = "xiaohong"
	(*p2).age = 21
	fmt.Println(*p2)
	// Go语言为了方便开发者访问结构体指针的成员变量，使用了语法糖（Syntactic sugar）技术，将 p2.name 形式转换为 (*p2).name
	p2.name = "xiaodong"
	fmt.Println(*p2)
	// *main.person, &{ 0}
	// {xiaohong 21}
	// {xiaodong 21}

	// 对结构体进行&取地址操作时，视为对该类型进行一次 new 的实例化操作
	p3 := &person{}
	(*p3).name = "xiaofei"
	(*p3).age = 22
	fmt.Println(*p3)
	// {xiaofei 22}

	// 取地址实例化是最广泛的一种结构体实例化方式，可以使用函数封装
	var version = 1
	cmd := newCommand(
		"version",
		&version,
		"show version",
	)
	fmt.Println(*cmd)
	// {version 0xc0000100a8 show version}

	// 使用“键值对”初始化结构体，，键值对形式的初始化适合选择性填充字段较多的结构体
	var p4 = person{
		name: "xiaotian",
		age:  27,
	}
	p5 := person{
		name: "xiaotian",
		// age:  27,
	}
	fmt.Println(p4, p5)
	// {xiaotian 27} {xiaotian 0}

	// Go语言可以在“键值对”初始化的基础上忽略“键”，也就是说，可以使用多个值的列表初始化结构体的字段。
	p6 := person{"xiaoli", 30}
	fmt.Println(p6)
	// {xiaoli 30}
	// 使用这种格式初始化时，需要注意：
	// 必须初始化结构体的所有字段。
	// 每一个初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	// 键值对与值列表的初始化形式不能混用。

	// 匿名结构体，没有类型名称，无须通过 type 关键字定义就可以直接使用
	dog := struct {
		name string
		ty   string
	}{
		// name: "wangwang",
		// ty:   "hashiqi",
		// 或者 直接：
		"wangwang", "hashiqi",
	}
	fmt.Println(dog)
	// {wangwang hashiqi}
}
