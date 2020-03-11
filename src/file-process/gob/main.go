package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

// Gob 是Go语言自己以二进制形式序列化和反序列化程序数据的格式，可以在 encoding 包中找到。这种格式的数据简称为 Gob（即 Go binary 的缩写）
// Gob 特定的用于纯 Go 的环境中，例如两个用Go语言写的服务之间的通信。这样的话服务可以被实现得更加高效和优化。

// 创建 gob 文件
func demo1() {
	info := map[string]string{
		"name":    "C语言中文网",
		"website": "http://c.biancheng.net/golang/",
	}

	File, _ := os.OpenFile("demo.gob", os.O_RDWR|os.O_CREATE, 0777) // 打开文件
	defer File.Close()

	enc := gob.NewEncoder(File) // 返回一个内容将写入 File 的编码器

	if err := enc.Encode(info); err != nil {
		fmt.Println(err)
	}
}

// 读取 gob 文件
func demo2() {
	var M map[string]string
	File, _ := os.Open("demo.gob")
	D := gob.NewDecoder(File)
	D.Decode(&M)
	fmt.Println(M)
}

func main() {
	demo1()

	demo2()
	// map[name:C语言中文网 website:http://c.biancheng.net/golang/]
}
