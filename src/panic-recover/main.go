package main

import "fmt"

// Recover 是一个Go语言的内建函数，可以让进入宕机流程中的 goroutine 恢复过来，recover 仅在延迟函数 defer 中有效，在正常的执行过程中，调用 recover 会返回 nil 并且没有其他任何效果，
// 如果当前的 goroutine 陷入恐慌，调用 recover 可以捕获到 panic 的输入值，并且恢复正常的执行。

func main() {
	fmt.Printf("hello world my name is %s, I'm %d\r\n", "songxingzhu", 26)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出了错：", err)
		}
	}()
	// 出了错： runtime error: integer divide by zero

	myPainc()
	fmt.Printf("这里应该执行不到！")
}

func myPainc() {
	var x = 30
	var y = 0
	// panic("我就是一个大错误！")
	var c = x / y // 此处引发 painc
	fmt.Println(c)
}
