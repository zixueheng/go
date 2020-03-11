package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

// Website 结构体
type Website struct {
	Name   string `xml:"name,attr"` // 这里的标记意思 将Name作为xml 标签 Website 的属性，并且使用小写 name 作为属性名
	Url    string
	Course []string
}

// xml 包还支持更为复杂的标签，包括嵌套。例如标签名为 'xml:"Books>Author"' 产生的是 <Books><Author>content</Author></Books> 这样的 XML 内容。同时除了 'xml:", attr"' 之外，
// 该包还支持 'xml:",chardata"' 这样的标签表示将该字段当做字符数据来写，支持 'xml:",innerxml"' 这样的标签表示按照字面量来写该字段，以及 'xml:",comment"' 这样的标签表示将该字段当做 XML 注释。
// 因此，通过使用标签化的结构体，我们可以充分利用好这些方便的编码解码函数，同时合理控制如何读写 XML 数据。

// 使用 encoidng/xml 包将 xml 数据存储到文件中
func demo1() {
	//实例化对象
	info := Website{"C语言中文网", "http://c.biancheng.net/golang/", []string{"Go语言入门教程", "Golang入门教程"}}

	f, err := os.Create("./info.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//序列化到文件中
	encoder := xml.NewEncoder(f)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("编码错误：", err.Error())
		return
	} else {
		fmt.Println("编码成功")
	}
}

// 从文件中读取 xml 并解析
func demo2() {
	//打开xml文件
	file, err := os.Open("./info.xml")
	if err != nil {
		// fmt.Printf("文件打开失败：%v", err)
		log.Fatal(err)
	}
	defer file.Close()

	info := Website{}

	//创建 xml 解码器
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Printf("解码失败：%v", err)
		return
	} else {
		fmt.Println("解码成功")
		fmt.Println(info)
	}
}

func main() {
	demo1()

	demo2()
	// {C语言中文网 http://c.biancheng.net/golang/ [Go语言入门教程 Golang入门教程]}
}
