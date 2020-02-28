package main

import (
	"fmt"
	"net/http"
	"sync"
)

// WaitGroup 对象内部有一个计数器，最初从0开始，它有三个方法：Add(), Done(), Wait() 用来控制计数器的数量。
// Add(n) 把计数器设置为n ，Done() 每次把计数器-1 ，wait() 会阻塞代码的运行，直到计数器地值减为0。

func main() {

	// 声明一个等待组
	var wg sync.WaitGroup

	// 准备一系列的网站地址
	var urls = []string{
		"http://www.github.com/",
		"https://www.qiniu.com/",
		"https://www.golangtc.com/",
	}

	// 遍历这些地址
	for _, url := range urls {

		// 每一个任务开始时，将等待组增加1
		wg.Add(1)

		// 开启一个并发
		go func(url string) {

			// 使用defer，表示函数完成时将等待组值减1
			defer wg.Done()

			// 使用http访问提供的地址
			_, err := http.Get(url)

			// 访问完成后，打印地址和可能发生的错误
			fmt.Println(url, err)

			// 通过参数传递url地址
		}(url)
	}

	// 等待所有的任务完成
	wg.Wait()

	fmt.Println("over")
}
