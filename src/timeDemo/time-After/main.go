package main

import (
	"fmt"
	"time"
)

// 调用time.After(duration)，此函数马上返回，返回一个time.Time类型的Chan，不阻塞。
// 后面你该做什么做什么，不影响。到了duration时间后，自动塞一个当前时间进去。
// 你可以阻塞的等待，或者晚点再取。
func demo1() {
	tchan := time.After(3 * time.Second)

	fmt.Printf("tchan type=%T\n", tchan)
	fmt.Println("mark 1")
	fmt.Println("tchan=", <-tchan) // 这里会阻塞等待数据
	fmt.Println("mark 2")
}

// time.AfterFunc() 函数是在 time.After 基础上增加了到时的回调，方便使用。
func demo2() {
	exit := make(chan int)
	fmt.Println("开始了")

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("3秒后")
		exit <- 1
	})

	<-exit // 阻塞等待
	fmt.Println("结束了")
}

// time.AfterFunc() 函数是在 time.After 基础上增加了到时的回调，方便使用。
// 而 time.After() 函数又是在 time.NewTimer() 函数上进行的封装

func main() {
	// demo1()
	// 先瞬间打印出前两行，然后等待3S，打印后后两行。
	// tchan type=<-chan time.Time
	// mark 1
	// tchan= 2020-03-02 16:18:39.4673768 +0800 CST m=+3.016099001
	// mark 2

	demo2()
	// 开始了
	// 3秒后
	// 结束了
}
