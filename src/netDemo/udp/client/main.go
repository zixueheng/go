package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 命令行输入 go run main.go 127.0.0.1:1200
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	checkError(err)

	conn, err := net.DialUDP("udp", nil, udpAddr) // 拨号连接
	checkError(err)

	_, err = conn.Write([]byte("anything")) // 写一些东西，但是服务端没有接受处理
	checkError(err)

	var buf [512]byte
	n, err := conn.Read(buf[0:]) // 从 连接中 读数据 放到 buf 数组中
	checkError(err)

	fmt.Println(string(buf[0:n]))
	os.Exit(0)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		os.Exit(1)
	}
}