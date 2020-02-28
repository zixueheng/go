package main

import "fmt"

// printer 打印函数
func printer(c chan int) {
	// 循环接收数据
	for {
		data := <-c // 接受数据并赋值给 data
		if data == 0 {
			break // 当 data 为0时认为接收结束
		}
		fmt.Printf("Print data: %d\n", data)
	}

	c <- 0 // 发送0通知 main 函数打印结束
}

func main() {
	c := make(chan int)
	go printer(c) // 开启并发(要在发送前就开启并发)
	for i := 1; i <= 10; i++ {
		fmt.Printf("Send data: %d\n", i)
		c <- i // 发送要打印的数据
	}

	c <- 0 // 发送数据 0 表示打印结束
	fmt.Println("Wait printer finished...")
	<-c // 接受 printer 发送的数据
	fmt.Println("Ok, Printer finished")
}
