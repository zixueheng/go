package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 示例1：
func running() {
	var times int
	// 构建一个无限循环
	for {
		times++
		fmt.Println("tick", times)
		// 延时1秒
		time.Sleep(time.Second)
	}
}

// 代码执行后，命令行会不断地输出 tick，同时可以使用 fmt.Scanln() 接受用户输入。两个环节可以同时进行。
// Go 程序在启动时，运行时（runtime）会默认为 main() 函数创建一个 goroutine。在 main() 函数的 goroutine 中执行到 go running 语句时，
// 归属于 running() 函数的 goroutine 被创建，running() 函数开始在自己的 goroutine 中执行。此时，main() 继续执行，两个 goroutine 通过 Go 程序的调度机制同时运作。
func demo1() {
	// 并发执行程序
	go running()
	// 接受命令行输入, 不做任何事情
	var input string
	fmt.Scanln(&input)
}

// 示例2：
var counter = 0

// 执行将counter累加，并打印的操作
func count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

// 使用加锁的方式进行变量在线程中的共享（这种方式是繁琐的）
func demo2() {
	lock := &sync.Mutex{}      // 初始化一个锁
	for i := 1; i <= 10; i++ { // 创建10个goroutine线程
		go count(lock)
	}
	// 使用 for 循环来不断检查 counter 的值（同样需要加锁）
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched() // 交出执行权，这个看不出来有什么作用
		if c >= 10 {      // 当其值达到 10 时，说明所有 goroutine 都执行完毕了，这时主函数返回，程序退出。
			break
		}
	}
}

func main() {
	// demo1()
	// demo2()

	// fmt.Println(runtime.NumCPU()) // 获取CPU核数
	// runtime.GOMAXPROCS(4) //设置需要用到的cpu数量
}
