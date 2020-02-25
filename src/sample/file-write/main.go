package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// file, err := os.OpenFile("./readme.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) // APPEND 表示追加到文件末尾
	file, err := os.OpenFile("./readme.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666) // TRUNC表示清空后再写
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	file.Write([]byte("hello golang\n"))
	file.WriteString("你好！\n")
	file.WriteString("\n")

	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	writer.WriteString("使用bufio写入的！\n")
	// 因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	// 所以要调用 flush方法，将缓存的数据真正写入到文件中。
	writer.Flush()

	file.Close()

	// 简便的写法，这里要注意每次写入会先清空文件
	err = ioutil.WriteFile("./readme1.txt", []byte("使用ioutil写入的！\n"), 0666)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
