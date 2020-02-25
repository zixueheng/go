package main

// 先要下载这个包：go get github.com/skip2/go-qrcode
import "github.com/skip2/go-qrcode"

func main() {
	qrcode.WriteFile("http://c.biancheng.net/", qrcode.Medium, 256, "./golang_qrcode.png")
}
