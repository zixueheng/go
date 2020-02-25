package main

import "fmt"

// 空接口可以保存任何类型这个特性可以方便地用于容器的设计。
// 下面例子使用 map 和 interface{} 实现了一个字典。字典在其他语言中的功能和 map 类似，可以将任意类型的值做成键值对保存，然后进行找回、遍历操作.

// Dictionary 字典
type Dictionary struct {
	data map[interface{}]interface{} // 键值都为interface{}类型
}

// Get 根据key键 取值
func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

// Set 设置键值对
func (d *Dictionary) Set(key interface{}, value interface{}) {
	d.data[key] = value
}

// Clear 清空字典
func (d *Dictionary) Clear() {
	// map 没有独立的复位内部元素的操作，需要复位元素时，使用 make 创建新的实例
	d.data = make(map[interface{}]interface{})
}

// NewDictionary 创建一个字典
func NewDictionary() *Dictionary {
	p := &Dictionary{}
	// 调用 Clear() 方法避免了 map 初始化过程的代码重复问题
	p.Clear()
	return p
}

// Visit 使用指定函数 遍历字典
// 参数是个回调函数，类型为 func(k,v interface{})bool，意思是返回键值数据（k、v）。bool 表示遍历流程控制，返回 true 时继续遍历，返回 false 时终止遍历
func (d *Dictionary) Visit(callback func(key, value interface{}) bool) {
	if callback == nil {
		return
	}
	// 对字典中的键值对 执行回调函数，如果返回 false 就结束遍历
	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}

func main() {
	dict := NewDictionary()
	dict.Set("My Factory", 60)
	dict.Set("Terra Craft", 36)
	dict.Set("Don't Hungry", 24)

	// 获取值及打印值
	favorite := dict.Get("Terra Craft")
	fmt.Println("favorite:", favorite)
	// favorite: 36

	dict.Visit(func(k, v interface{}) bool {
		// 使用类型断言将 v 转成int
		if v.(int) > 40 {
			fmt.Println(k, v, "很贵")
			return true
		}
		fmt.Println(k, v, "很便宜")
		return true
	})
	// My Factory 60 很贵
	// Terra Craft 36 很便宜
	// Don't Hungry 24 很便宜
}
