package base

import "fmt"

// 工程基类，主要注册类和创建类

// Class 类接口
type Class interface {
	// 规定一个 String 方法 返回类名
	String() string
}

// factoryByName 是一个 map 键是类名称，值是返回Class的函数
var factoryByName = make(map[string]func() Class)

// Register 根据类名注册类
func Register(className string, f func() Class) {
	factoryByName[className] = f
}

// Create 创建类
func Create(className string) Class {
	f, ok := factoryByName[className]
	if ok {
		return f()
	}
	fmt.Printf("Class: %s not found\n", className)
	return nil
}
