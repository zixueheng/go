package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("hello golang")

	// 向终端中输出文本
	fmt.Fprintln(os.Stdout, "hello golang") //

	// 打开一个文件
	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error: %T", err)
		return
	}
	defer file.Close() // 关闭文件的代码要放在“错误判断”后面，因为 如果发生错误，file变量为nil调用 Close()发放会报错
	// 向文件中写入文本
	fmt.Fprintln(file, "hello golang") // file 实现了 io.Wrinter 接口，所以能传给 Fprintln()
}
