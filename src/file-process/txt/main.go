package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 写入纯文本
func demo1() {
	//创建一个新文件，写入内容
	filePath := "./output.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("打开文件错误= %v \n", err)
		return
	}
	defer file.Close() //及时关闭

	//写入内容
	str := "http://c.biancheng.net/golang/\n" // \n\r表示换行  txt文件要看到换行效果要用 \r\n
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		writer.WriteString(str)
	}

	//因为 writer 是带缓存的，因此在调用 WriterString 方法时，内容是先写入缓存的
	//所以要调用 flush方法，将缓存的数据真正写入到文件中。
	writer.Flush()
}

// 读纯文本
func demo2() {
	// 打开文件
	file, err := os.Open("./output.txt")
	if err != nil {
		fmt.Println("文件打开失败 = ", err)
	}
	defer file.Close() // 及时关闭 file 句柄，否则会有内存泄漏

	// 创建一个 *Reader ， 是带缓冲的
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // 读到一个换行就结束
		if err == io.EOF {                  // io.EOF 表示文件的末尾
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}

func main() {
	demo1()
	demo2()
}
