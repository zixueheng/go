package main

import "fmt"

// 单向通道
// 格式：
// var 通道实例 chan<- 元素类型    // 只能发送通道
// var 通道实例 <-chan 元素类型    // 只能接收通道

func main() {
	// ch := make(chan int)
	var ch chan int = make(chan int)
	var chSend chan<- int = ch // 箭头在后 是 发送
	var chGet <-chan int = ch  // 箭头在前 是 接受
	go getData(chGet)          // 开启并发(要在发送前就开启并发)
	sendData(chSend)

	chSend <- 0 // 发送数据 0 表示打印结束
}

// getData函数，参数是单向接收通道
func getData(chGet <-chan int) {
	// 循环接收数据
	for {
		data := <-chGet // 接受数据并赋值给 data
		if data == 0 {
			break // 当 data 为0时认为接收结束
		}
		fmt.Printf("Get data: %d\n", data)
	}
	// chGet <- 100 // 发送一个100 试试：报错（send to receive-only type <-chan int）意思chGet只能接收数据
}

// sendData函数，参数是单向发送通道
func sendData(chSend chan<- int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("Send data: %d\n", i)
		chSend <- i // 发送要打印的数据
	}
}
