package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读取文件示例，只能读取小文档，读取大文件可能会造成内存的大量占用

// ReadFileContent1 普通的读取文件内容
func ReadFileContent1(filename string) (string, error) {
	file, err := os.Open(filename) // os.Open() 打开的文件是只读的
	if err != nil {
		return "", fmt.Errorf("open file failed: %v", err)
	}
	defer file.Close() // 延迟关闭文件

	var content string
	var tmp [128]byte // 声明一个 126 字节的数组
	for {
		n, e := file.Read(tmp[:]) // 读到数据放到 tmp切片当中，这里 tmp 切片是从上面128位的数组切出来的
		if e != nil {
			return "", fmt.Errorf("read file error: %v", err)
		}
		content += string(tmp[:n]) // 从 tmp 开头切到实际读取的长度
		if n < 128 {
			break
		}
	}
	return content, nil
}

// ReadFileContent2 使用 bufio 读取文件示例
func ReadFileContent2(filename string) (string, error) {
	file, err := os.Open(filename) // os.Open() 打开的文件是只读的
	if err != nil {
		return "", fmt.Errorf("open file failed: %v", err)
	}
	defer file.Close() // 延迟关闭文件

	var content string
	reader := bufio.NewReader(file)
	for {
		s, e := reader.ReadString('\n') //  \n 换行符，是个字符 要用单引号
		if e == io.EOF {                // io.EOF 表示读到文件末尾，必须要放在下面的报错前面
			break
		}
		if e != nil {
			return "", fmt.Errorf("read file error: %v", e)
		}
		content += s
	}

	return content, nil
}
func main() {
	// content, err := ReadFileContent1("./main.go")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(content)
	// }

	// content, err := ReadFileContent2("./main.go")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(content)
	// }

	// 最简单的读文件方式
	bytes, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(bytes))
	}
}
