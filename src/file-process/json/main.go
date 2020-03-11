package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Website 结构体
type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

func demo1() {
	// info := []Website{{"Golang", "http://c.biancheng.net/golang/", []string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"}}, {"Java", "http://c.biancheng.net/java/", []string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}}}

	info := []Website{
		Website{"Golang1", "http://c.biancheng.net/golang/", []string{"http://c.biancheng.net/cplus/", "http://c.biancheng.net/linux_tutorial/"}},
		Website{"Java1", "http://c.biancheng.net/java/", []string{"http://c.biancheng.net/socket/", "http://c.biancheng.net/python/"}},
	}

	file, err := os.Create("test.json") // 创建一个 新文件，如果已存在会清空源文件
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file) // 创建已个json编码器，编码的内容写入 file 中
	err = encoder.Encode(info)       // Encode将info的json编码写入输出流file中，并会写入一个换行符

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("编码成功")
	}
}

// 从文件中读取json解析成结构体数据
func demo2() {
	file, err := os.Open("test.json") // 返回只读的 file 指针
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var info []Website
	decoder := json.NewDecoder(file) // 创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	err = decoder.Decode(&info)      // 从输入流 file 读取下一个json编码值并保存在 file 指向的值里
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(info)
	}
}

func main() {
	demo1()

	demo2()
	// [{Golang http://c.biancheng.net/golang/ [http://c.biancheng.net/cplus/ http://c.biancheng.net/linux_tutorial/]} {Java http://c.biancheng.net/java/ [http://c.biancheng.net/socket/ http://c.biancheng.net/python/]}]
}
