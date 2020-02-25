package main

import (
	"bufio"
	"fmt"
	"os"
)

// ScanInput 从终端中获取输入
// 这里会有一个问题，当输入的是连续字符输出是完整的：
// 请输入内容：kdkajdsajd
// 你输入的是：kdkajdsajd
// 当输入中间带空格只能输出第一个空格前的内容：
// 请输入内容：a hm mcd
// 你输入的是：a
func ScanInput() {
	fmt.Print("请输入内容：")
	var str string
	fmt.Scanln(&str)
	fmt.Printf("你输入的是：%s\n", str)
}

// ScanInputByBufio 使用 bufio 从终端 os.Stdin 读取内容
// 请输入内容：a v d
// 你输入的是：a v d
func ScanInputByBufio() {
	fmt.Print("请输入内容：")
	var str string
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n') // 从 reader 中读取 字符串 到 换行(\n)截止
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("你输入的是：%s\n", str)
}

func main() {
	// ScanInput()
	ScanInputByBufio()
}
