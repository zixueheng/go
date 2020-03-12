package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// buffer 是缓冲器的意思，Go语言要实现缓冲读取需要使用到 bufio 包。
// bufio 包本身包装了 io.Reader 和 io.Writer 对象，同时创建了另外的 Reader 和 Writer 对象，因此对于文本 I/O 来说，bufio 包提供了一定的便利性。

// buffer 缓冲器的实现原理就是，将文件读取进缓冲（内存）之中，再次读取的时候就可以避免文件系统的 I/O 从而提高速度。同理在进行写操作时，先把文件写入缓冲（内存），然后由缓冲写入文件系统。

// bufio 和 io 包中有很多操作都是相似的，唯一不同的地方是 bufio 提供了一些缓冲的操作，如果对文件 I/O 操作比较频繁的，使用 bufio 包能够提高一定的性能。

// 使用 bufio 包写入文件
func demo1() {
	name := "demo.txt"
	content := "http://c.biancheng.net/golang/"
	fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer fileObj.Close()

	writeObj := bufio.NewWriterSize(fileObj, 4096)
	//使用 Write 方法,需要使用 Writer 对象的 Flush 方法将 buffer 中的数据刷到磁盘
	buf := []byte(content)
	if _, err := writeObj.Write(buf); err == nil {
		if err := writeObj.Flush(); err != nil {
			panic(err)
		}
		fmt.Println("数据写入成功")
	}
}

// 使用 bufio 包读取文件
func demo2() {
	fileObj, err := os.Open("demo.txt")
	if err != nil {
		fmt.Println("文件打开失败：", err)
		return
	}
	defer fileObj.Close()

	// 一个文件对象本身是实现了io.Reader的 使用bufio.NewReader去初始化一个Reader对象，存在buffer中的，读取一次就会被清空
	reader := bufio.NewReader(fileObj)
	buf := make([]byte, 1024)
	// 读取 Reader 对象中的内容到 []byte 类型的 buf 中
	info, err := reader.Read(buf)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("读取的字节数:" + strconv.Itoa(info))
	// 这里的buf是一个[]byte，因此如果需要只输出内容，仍然需要将文件内容的换行符替换掉
	fmt.Println("读取的文件内容:", string(buf))
}
func main() {
	demo1()

	demo2()
}
