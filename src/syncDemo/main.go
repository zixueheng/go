package main

import (
	"fmt"
	"sync"
	"time"
)

// Go语言中 sync 包里提供了互斥锁 Mutex 和读写锁 RWMutex 用于处理并发过程中可能出现同时两个或多个协程（或线程）读或写同一个变量的情况。

func main() {
	var a = 0
	var lock sync.Mutex // 声明一个 互斥锁 Mutex
	for i := 0; i < 10; i++ {
		go func(idx int) {
			lock.Lock()         // 一个协程开始执行的时候持有锁
			defer lock.Unlock() // 协程结束的时候释放锁对象
			a++
			fmt.Printf("goroutine %d, a=%d\n", idx, a)
		}(i)
	}
	// 等待 1s 结束主程序
	// 确保所有协程执行完，再退出程序，如果不加看不到输出
	time.Sleep(time.Second)
	// goroutine 0, a=1
	// goroutine 5, a=2
	// goroutine 4, a=3
	// goroutine 7, a=4
	// goroutine 8, a=5
	// goroutine 9, a=6
	// goroutine 1, a=7
	// goroutine 3, a=8
	// goroutine 6, a=9
	// goroutine 2, a=10
	// 从上面输出的顺序看，协程不是按顺序执行的
}
