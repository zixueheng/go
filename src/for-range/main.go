package main

import "fmt"

// for range 可以遍历数组、切片、字符串、map 及通道（channel），for range 语法上类似于其它语言中的 foreach 语句，一般形式为：
// for key, val := range coll {
//     ...
// }
// val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值。

func main() {
	// 遍历数组、切片
	for key, value := range []int{1, 2, 3, 4} {
		fmt.Printf("key:%d  value:%d\n", key, value)
	}
	// key:0  value:1
	// key:1  value:2
	// key:2  value:3
	// key:3  value:4

	// 遍历字符串
	var str = "hello 你好"
	for key, value := range str {
		fmt.Printf("key:%d value:0x%x\n", key, value)
	}
	// key:0 value:0x68
	// key:1 value:0x65
	// key:2 value:0x6c
	// key:3 value:0x6c
	// key:4 value:0x6f
	// key:5 value:0x20
	// key:6 value:0x4f60
	// key:9 value:0x597d
	// 中文占三个字节，所以 key 6 下面是 9

	// 遍历 map
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}

	// 遍历通道（channel）——接收通道数据
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
	// 1
	// 2
	// 3

	// Go语言的 for 包含初始化语句、条件表达式、结束语句，这 3 个部分均可缺省。
	// for range 支持对数组、切片、字符串、map、通道进行遍历操作。
	// 在需要时，可以使用匿名变量对 for range 的变量进行选取。

}
