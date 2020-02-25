package main

import (
	"fmt"
	"os"
)

// Go语言的 os 包中提供了操作系统函数的接口，是一个比较重要的包。
// 顾名思义，os 包的作用主要是在服务器上进行系统的基本操作，如文件操作、目录操作、执行命令、信号与中断、进程、系统状态等等。

func main() {
	hostname, _ := os.Hostname()
	fmt.Println(hostname) // 主机名

	envs := os.Environ() // 所有的环境变量，返回值格式为“key=value”的字符串的切片拷贝
	fmt.Println(envs)
	fmt.Println()
	// for i, v := range envs {
	// 	fmt.Println(i, v)
	// }

	path := os.Getenv("PATH") // 系统的 PATH 配置
	fmt.Println(path)

	wd, _ := os.Getwd() // 对应当前工作目录的根路径
	fmt.Println(wd)     // /Users/heyongliang/go/src/osDemo

	os.Mkdir("temp", 0755) // 在当前目录创建一个 temp 目录

	os.MkdirAll("./1/2/3", 0755) // 递归的创建目录

	os.Remove("temp") // 删除当前目录下的 temp

	os.RemoveAll("./1") // 递归的删除 ./1 下所有目录 包括文件
}
