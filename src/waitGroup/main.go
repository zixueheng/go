package main

import (
	"fmt"
	"net/http"
	"sync"
)

// WaitGroup 对象内部有一个计数器，最初从0开始，它有三个方法：Add(), Done(), Wait() 用来控制计数器的数量。
// Add(n) 把计数器设置为n ，Done() 每次把计数器-1 ，wait() 会阻塞代码的运行，直到计数器地值减为0。

// 等待组内部拥有一个计数器，计数器的值可以通过方法调用实现计数器的增加和减少。当我们添加了 N 个并发任务进行工作时，就将等待组的计数器值增加 N。每个任务完成时，这个值减 1。
// 同时，在另外一个 goroutine 中等待这个等待组的计数器值为 0 时，表示所有任务已经完成。
func main() {

	// 声明一个等待组，对一组等待任务只需要一个等待组，而不需要每一个任务都使用一个等待组
	var wg sync.WaitGroup

	// 网站地址 的字符串切片
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
