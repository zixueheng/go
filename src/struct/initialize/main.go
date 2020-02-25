package main

import "fmt"

// 结构体内嵌初始化时，将结构体内嵌的类型作为字段名像普通结构体一样进行初始化

// Wheel 车轮
type Wheel struct {
	Size int
}

// Engine 引擎
type Engine struct {
	Power int    // 功率
	Type  string // 类型
}

// Car 车
type Car struct {
	Wheel
	Engine
}

// Car2 车
// 有时考虑编写代码的便利性，会将结构体直接定义在嵌入的结构体中。也就是说，结构体的定义不会被外部引用到。在初始化这个被嵌入的结构体时，就需要再次声明结构才能赋予数据
type Car2 struct {
	Wheel
	Engine struct {
		Power int    // 功率
		Type  string // 类型
	}
}

func main() {
	car := Car{
		Wheel: Wheel{
			Size: 16,
		},
		Engine: Engine{
			Power: 160,
			Type:  "2.0T",
		},
	}

	fmt.Printf("%+v\n", car)
	// {Wheel:{Size:16} Engine:{Power:160 Type:2.0T}}

	car2 := Car2{
		Wheel: Wheel{
			Size: 18,
		},
		// 由于 Engine 字段的类型并没有被单独定义，因此在初始化其字段时需要先填写 struct{…} 声明其类型：
		Engine: struct {
			Power int
			Type  string
		}{
			Power: 180,
			Type:  "3.0T",
		},
	}
	fmt.Printf("%+v\n", car2)
	// {Wheel:{Size:18} Engine:{Power:180 Type:3.0T}}
}
