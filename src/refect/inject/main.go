package main

import (
	"fmt"

	"github.com/codegangsta/inject"
)

type s1 interface{}
type s2 interface{}

func format(name string, company s1, level s2, age int, wgt int) {
	fmt.Printf("name ＝ %s, company=%s, level=%s, age ＝ %d, wgt = %d!\n", name, company, level, age, wgt)
}

// inject 包借助反射实现函数的注入调用
func demo1() {
	//控制实例的创建
	inj := inject.New()
	//实参注入
	inj.Map("tom")
	inj.MapTo("tencent", (*s1)(nil))
	inj.MapTo("T4", (*s2)(nil))
	inj.Map(23)
	inj.Map(60) // 这里会出现问题，传递的60 会覆盖上面的 23
	// 函数反转调用
	inj.Invoke(format)
}

// inject 包不但提供了对函数的注入，还实现了对 struct 类型的注入
func demo2() {
	//创建被注入实例
	s := Staff{}

	//控制实例的创建
	inj := inject.New()

	// 初始化注入值
	inj.Map("tom")
	inj.MapTo("tencent", (*s1)(nil))
	inj.MapTo("T4", (*s2)(nil))
	inj.Map(23)

	// Apply 方法是用于对 struct 的字段进行注入，参数为指向底层类型为结构体的指针。可注入的前提是：字段必须是导出的（也即字段名以大写字母开头），并且此字段的 tag 设置为`inject`
	inj.Apply(&s)

	//打印结果
	fmt.Printf("s ＝ %v\n", s)
}

// Staff 结构体
type Staff struct {
	Name    string `inject`
	Company s1     `inject`
	Level   s2     `inject`
	Age     int    `inject`
}

func main() {
	demo1()
	// name ＝ tom, company=tencent, level=T4, age ＝ 23!

	demo2()
	// s ＝ {tom tencent T4 23}
}
