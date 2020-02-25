package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// 用Go语言搭建一个 Web 服务器

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	// 运行之后并没有什么提示信息，但是命令行窗口会被占用（不能再输入其它命令）。这时我们在浏览器中输入 localhost:8000
}

func index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "C语言中文网")
	content, _ := ioutil.ReadFile("index.html")
	w.Write(content)
}
