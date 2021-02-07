package main

import (
	"fmt"
	"reflect"
)

// Cater 接口
type Cater interface {
	Eat()
}

// Cat ...
type Cat struct {
	Name string
}

// Eat 实现 Cater 接口
func (c *Cat) Eat() {
	fmt.Println(c.Name + " 吃吃")
}

func main() {
	var C Cater = &Cat{Name: "Tom"}
	C.Eat()
	// Tom 吃吃

	var typeOfC reflect.Type = reflect.TypeOf(C)
	fmt.Println(typeOfC.Name(), typeOfC.Kind())
	// "" ptr
	if typeOfC.Kind() == reflect.Ptr {
		typeOfC = typeOfC.Elem() // 取出指针类型的实际反射类型
		fmt.Println(typeOfC.Name(), typeOfC.Kind())
		// Cat struct
	}

	var valueOfC reflect.Value = reflect.ValueOf(C)
	fmt.Println(valueOfC.Kind(), valueOfC.Interface()) // 返回v 的种类 和 v接口类型值
	// ptr &{Tom}
	if valueOfC.Kind() == reflect.Ptr {
		v := valueOfC.Elem()
		fmt.Println(v.Kind(), v.Interface()) // 返回v 的种类 和 v接口类型值
		// struct {Tom}

		var anther Cat = v.Interface().(Cat) // 将 v接口类型值 断言成Cat类型
		fmt.Println(anther)
		// {Tom}
	}
}
