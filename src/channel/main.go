package main

import "fmt"
import "time"

func main() {
	// demo1()
	// 1 wait goroutine
	// 2 start goroutine
	// 3 exit goroutine
	// 4 all done

	demo2()
	// Start send num: 3
	// End send num: 3
	// Get num: 3
	// Start send num: 2
	// End send num: 2
	// Get num: 2
	// Start send num: 1
	// End send num: 1
	// Get num: 1
	// Start send num: 0
	// End send num: 0
	// Get num: 0
}

// 使用通道做并发同步
func demo1() {
	// 构建一个同步用的通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		fmt.Println("2 start goroutine")
		// 匿名 goroutine 即将结束时，通过通道通知 demo1 的 goroutine，这一句会一直阻塞直到 demo1 的 goroutine 接收为止
		ch <- 0
		fmt.Println("3 exit goroutine")
	}()
	fmt.Println("1 wait goroutine")
	// 等待匿名goroutine
	<-ch // 执行该语句时将会发生阻塞，直到接收到数据，但接收到的数据会被忽略。这个方式实际上只是通过通道在 goroutine 间阻塞收发实现并发同步
	// data := <-ch // 执行该语句时将会阻塞，直到接收到数据并赋值给 data 变量
	// fmt.Println(data) //打印获取的 值：0
	fmt.Println("4 all done")
}

// 使用 for 从通道中接收数据
func demo2() {
	// 构建一个通道
	ch := make(chan int)
	// 开启一个并发匿名函数
	go func() {
		// 从3循环到0
		for i := 3; i >= 0; i-- {
			fmt.Printf("Start send num: %d\n", i)
			// 发送3到0之间的数值
			ch <- i
			fmt.Printf("End send num: %d\n", i)
			// 每次发送后暂停 3 秒
			time.Sleep(3 * time.Second)
		}
	}()
	// 遍历接收通道数据，每次接收到数据时会等待3秒
	// 上一个循环接收通道数据 会阻塞下一个循环 直到它接收到数据
	for data := range ch {
		// 打印通道数据
		fmt.Printf("Get num: %d\n", data)
		// 当遇到数据0时, 退出接收循环
		if data == 0 {
			break // 如果继续发送，由于接收 goroutine 已经退出，没有 goroutine 发送到通道，因此运行时将会触发宕机报错: fatal error: all goroutines are asleep - deadlock!
		}
	}
}
