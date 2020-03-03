package main

import "fmt"

// 通道是一个引用对象，和 map 类似。map 在没有任何外部引用时，
// Go语言程序在运行时（runtime）会自动对内存进行垃圾回收（Garbage Collection, GC）。类似的，通道也可以被垃圾回收，但是通道也可以被主动关闭。

// 给被关闭通道发送数据将会触发 panic
// 被关闭的通道不会被置为 nil。如果尝试对已经关闭的通道进行发送，将会触发宕机
func demo1() {
	// 创建一个整型的通道
	ch := make(chan int)
	// 关闭通道
	close(ch) // 关闭通道，注意 ch 不会被 close 设置为 nil，依然可以被访问
	// 打印通道的指针, 容量和长度
	fmt.Printf("ptr:%p cap:%d len:%d\n", ch, cap(ch), len(ch))
	// 给关闭的通道发送数据
	ch <- 1
}

// 从已经关闭的通道接收数据或者正在接收数据时，将会接收到通道类型的零值，然后停止阻塞并返回。
func demo2() {
	// 创建一个整型带两个缓冲的通道
	ch := make(chan int, 2)

	// 给通道放入两个数据
	ch <- 11
	ch <- 12

	// 关闭缓冲
	close(ch)
	// 遍历缓冲所有数据, 且多遍历1个
	for i := 0; i < cap(ch)+1; i++ {
		// 从通道中取出数据
		v, ok := <-ch

		// 打印取出数据的状态
		fmt.Println(v, ok)
	}
}

func main() {
	// demo1()
	// ptr:0xc000048060 cap:0 len:0
	// panic: send on closed channel

	demo2() //运行结果前两行正确输出带缓冲通道的数据，表明缓冲通道在关闭后依然可以访问内部的数据。
	// 11 true
	// 12 true
	// 0 false // 运行结果第三行的“0 false”表示通道在关闭状态下取出的值。0 表示这个通道的默认值，false 表示没有获取成功，因为此时通道已经空了。我们发现，在通道关闭后，即便通道没有数据，在获取时也不会发生阻塞，但此时取出数据会失败。
}
