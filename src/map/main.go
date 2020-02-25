package main

import (
	"fmt"
	"sync"
)

// Go语言中 map 是一种特殊的数据结构，一种元素对（pair）的无序集合，pair 对应一个 key（索引）和一个 value（值），所以这个结构也称为关联数组或字典，
// 这是一种能够快速寻找值的理想结构，给定 key，就可以迅速找到对应的 value。
// map 是引用类型，可以使用如下方式声明：
// var mapname map[keytype]valuetype

func main() {
	var mapLit map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
	// 	Map literal at "one" is: 1
	// Map created at "key2" is: 3.141590
	// Map assigned at "two" is: 3
	// Map literal at "ten" is: 0

	scene := make(map[string]int)
	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	// 根据键名取值，ok 表示是否存在，v表示值，如果不存在 v 是对应值类型的零值（int零值0，string零值""）
	v, ok := scene["china"]
	fmt.Println(v, ok) // 960 true

	v, ok = scene["usa"]
	fmt.Println(v, ok) // 0 false

	// 遍历键和值
	for k, v := range scene {
		fmt.Println(k, v)
	}
	// 只遍历值
	for _, v := range scene {
		fmt.Println(v)
	}
	// 只遍历键名
	for k := range scene {
		fmt.Println(k)
	}
	// 删除某个键值
	delete(scene, "brazil")

	// map 在并发情况下，只读是线程安全的，同时读写是线程不安全的
	// sync.Map 和 map 不同，不是以语言原生形态提供，而是在 sync 包下的特殊结构
	var scene1 sync.Map
	// 将键值对保存到sync.Map
	scene1.Store("greece", 97)
	scene1.Store("london", 100)
	scene1.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene1.Load("london"))
	// 根据键删除对应的键值对
	scene1.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene1.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
	// sync.Map 有以下特性：
	// 无须初始化，直接声明即可。
	// sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。
	// 使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false。

}
