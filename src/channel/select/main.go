package main

import (
	"fmt"
	"time"
)

// channel的 超时使用 select 实现

// 调用time.After(duration)，此函数马上返回，返回一个time.Time类型的Chan，不阻塞。
// 后面你该做什么做什么，不影响。到了duration时间后，自动塞一个当前时间进去。
// 你可以阻塞的等待，或者晚点再取。

//
func main() {
	ch := make(chan int)
	quit := make(chan bool)

	//新开一个协程
	go func() {
		for {
			select {
			case num := <-ch: // 持续接收数据
				fmt.Println("num = ", num)
			case <-time.After(5 * time.Second): // 这里意思等待5秒后，进行超时操作（<-是在通道里面取值了，会阻塞执行）
				fmt.Println("超时")
				quit <- true // quit 里面放入 true
			}
		}
	}() // 别忘了()

	// 每隔1秒往 ch 里面放一个数
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-quit // 当quit 取到值时就意味着程序结束了
	fmt.Println("程序结束")
}
