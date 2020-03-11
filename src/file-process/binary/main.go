package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

// 虽然Go语言的 encoding/gob 包非常易用，而且使用时所需代码量也非常少，但是我们仍有可能需要创建自定义的二进制格式。自定义的二进制格式有可能做到最紧凑的数据表示，并且读写速度可以非常快。

// 不过，在实际使用中，我们发现以Go语言二进制格式的读写通常比自定义格式要快非常多，而且创建的文件也不会大很多。
// 但如果我们必须通过满足 gob.GobEncoder 和 gob.GobDecoder 接口来处理一些不可被 gob 编码的数据，这些优势就有可能会失去。

// 在有些情况下我们可能需要与一些使用自定义二进制格式的软件交互，因此了解如何处理二进制文件就非常有用。

type Website struct {
	Url int32
}

// 写自定义二进制文件
func demo1() {
	file, err := os.Create("output.bin") // 创建二进制文件
	if err != nil {
		fmt.Println("文件创建失败 ", err.Error())
		return
	}
	defer file.Close()

	for i := 1; i <= 10; i++ {
		info := Website{
			int32(i),
		}

		var binBuf bytes.Buffer

		// func binary.Write(w io.Writer, order ByteOrder, data interface{}) error
		// Write 函数可以将参数 data 的 binary 编码格式写入参数 w 中，参数 data 必须是定长值、定长值的切片、定长值的指针。
		// 参数 order 指定写入数据的字节序，写入结构体时，名字中有_的字段会置为 0
		binary.Write(&binBuf, binary.LittleEndian, info) // 将 info 的 二进制编码 写入 缓存binBuf 中
		b := binBuf.Bytes()                              // 获取缓存的 二进制 数据
		_, err = file.Write(b)                           // 将二进制数据b 写入 文件 file
		if err != nil {
			fmt.Println("编码失败", err.Error())
			return
		}
	}
	fmt.Println("编码成功")
}

// 读取二进制文件
func demo2() {
	file, err := os.Open("output.bin")
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}

	m := Website{}
	for i := 1; i <= 10; i++ {
		data := readNextBytes(file, 4)
		buffer := bytes.NewBuffer(data)

		err = binary.Read(buffer, binary.LittleEndian, &m) // 从缓存buffer中读取二进制数据放到 m 中
		if err != nil {
			fmt.Println("二进制文件读取失败", err)
			return
		}
		fmt.Println("第", i, "个值为：", m)
	}
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		fmt.Println("解码失败", err)
	}
	return bytes
}

func main() {
	demo1()

	demo2()
	// 第 1 个值为： {1}
	// 第 2 个值为： {2}
	// 第 3 个值为： {3}
	// 第 4 个值为： {4}
	// 第 5 个值为： {5}
	// 第 6 个值为： {6}
	// 第 7 个值为： {7}
	// 第 8 个值为： {8}
	// 第 9 个值为： {9}
	// 第 10 个值为： {10}
}
