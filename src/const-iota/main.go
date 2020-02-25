package main

import "fmt"

// 批量声明的常量，除了第一个外其它的常量右边的初始化表达式都可以省略，如果省略初始化表达式则表示使用前面常量的初始化表达式，对应的常量类型也是一样的
const (
	a = 1
	b
	c = 2
	d
)

// iota 常量生成器
// 常量声明可以使用 iota 常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式。
// 在一个 const 声明语句中，在第一个声明的常量所在的行，iota 将会被置为 0，然后在每一个有常量声明的行加一

// Weekday 定义为 int类型
type Weekday int

// 批量常量
const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(a, b, c, d)
	// 1 1 2 2

	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	// 0 1 2 3 4 5 6
}
