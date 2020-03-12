package main

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

// tar 是一种打包格式，但不对文件进行压缩，所以打包后的文档一般远远大于 zip 和 tar.gz，因为不需要压缩的原因，所以打包的速度是非常快的，打包时 CPU 占用率也很低。
// tar 的目的在于方便文件的管理

// tar 打包实现原理如下：
//     创建一个文件 x.tar，然后向 x.tar 写入 tar 头部信息；
//     打开要被 tar 的文件，向 x.tar 写入头部信息，然后向 x.tar 写入文件信息；
//     当有多个文件需要被 tar 时，重复第二步直到所有文件都被写入到 x.tar 中；
//     关闭 x.tar，完成打包。

// tar打包文件
func demo1() {
	f, err := os.Create("./output.tar") // 创建一个 tar 文件
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	tw := tar.NewWriter(f) // 创建一个 写入器 ，内容将被写入 f
	defer tw.Close()

	fileinfo, err := os.Stat("./readme.md") // 获取文件相关信息
	if err != nil {
		fmt.Println(err)
	}

	hdr, err := tar.FileInfoHeader(fileinfo, "") // FileInfoHeader返回一个根据 fi 填写了部分字段的 Header
	if err != nil {
		fmt.Println(err)
	}

	err = tw.WriteHeader(hdr) // 写入头文件信息
	if err != nil {
		fmt.Println(err)
	}

	f1, err := os.Open("./readme.md")
	if err != nil {
		fmt.Println(err)
		return
	}

	m, err := io.Copy(tw, f1) //将readme.md文件中的信息写入压缩包中
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)
}

// 解压 tar 归档文件比创建 tar 归档文档稍微简单些。
// 首先需要将其打开，然后从这个 tar 头部中循环读取存储在这个归档文件内的文件头信息，从这个文件头里读取文件名，以这个文件名创建文件，然后向这个文件里写入数据即可。
func demo2() {
	f, err := os.Open("output.tar")
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer f.Close()

	r := tar.NewReader(f)
	// 按文件遍历 tar
	for hdr, err := r.Next(); err != io.EOF; hdr, err = r.Next() { // Next() 转入tar档案文件下一记录，它会返回下一记录的头域
		if err != nil {
			fmt.Println(err)
			return
		}
		fileinfo := hdr.FileInfo()
		fmt.Println(fileinfo.Name())
		f, err := os.Create("new_" + fileinfo.Name()) // 创建新文件 以new_ 作为前缀 防止覆盖
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		_, err = io.Copy(f, r) // 将内容写入 新创建的文件中
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	demo1()

	demo2()
}
