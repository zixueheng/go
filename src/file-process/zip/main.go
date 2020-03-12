package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"os"
)

// Go语言提供了 archive/zip 包来操作压缩文件

// 创建 zip 归档文件
func demo1() {
	var buff bytes.Buffer                     // 声明缓存
	var zw *zip.Writer = zip.NewWriter(&buff) // 创建并返回一个将zip文件写入 buff 的*Writer

	// 结构体切片
	var files = []struct{ name, body string }{
		{"file1.txt", "aaaaaaaaaaaaaa"},
		{"file2.txt", "bbbbbbbbbbbbbb"},
		{"file3.txt", "cccccccccccccc"},
	}

	// 遍历切片
	for _, file := range files {
		// Create 方法使用给出的文件名添加一个文件进zip文件。本方法返回一个io.Writer接口（用于写入新添加文件的内容）。
		// 文件名必须是相对路径，不能以设备或斜杠开始，只接受'/'作为路径分隔。新增文件的内容必须在下一次调用CreateHeader、Create或Close方法之前全部写入
		w, err := zw.Create(file.name) // 将 文件名 添加到 zip 文件中，返回一个 写入器 等待写入内容，内容会以压缩的方式写入
		if err != nil {
			fmt.Println(err)
		}
		_, err = w.Write([]byte(file.body)) // 将内容 file.body 写入
		if err != nil {
			fmt.Println(err)
		}
	}

	// 写入完成后 要调用 Close() 关闭zip文件的的写入
	err := zw.Close()
	if err != nil {
		fmt.Println(err)
	}

	zipfile, _ := os.Create("file.zip") // 创建一个文件

	buff.WriteTo(zipfile) // 将缓存的内容写入 上面创建的 文件
}

// 读取 zip 归档文件
func demo2() {
	// OpenReader会打开name指定的zip文件并返回一个*ReadCloser
	rc, err := zip.OpenReader("file.zip") // 返回的 ReadCloser 结构体 包含 File切片（即内部压缩的文件）
	if err != nil {
		fmt.Println(err)
	}
	defer rc.Close() // Close关闭zip文件，使它不能用于I/O。

	// 遍历压缩文件中的所有文件
	for _, file := range rc.File {
		fmt.Printf("文件名: %s\n", file.Name)
		rc2, err := file.Open() // 打开文件
		if err != nil {
			fmt.Printf(err.Error())
		}

		_, err = io.CopyN(os.Stdout, rc2, int64(file.UncompressedSize64)) // 将文件内容拷贝到终端中显示
		if err != nil {
			fmt.Printf(err.Error())
		}
		fmt.Println()
		rc2.Close()
	}
}

func main() {
	demo1()

	demo2()
	// 文件名: file1.txt
	// aaaaaaaaaaaaaa
	// 文件名: file2.txt
	// bbbbbbbbbbbbbb
	// 文件名: file3.txt
	// cccccccccccccc
}
