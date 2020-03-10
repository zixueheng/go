package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// 生成一个 tcp 地址
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8000") // 端口号前面地址不写表示本机
	if err != nil {
		log.Fatal(err)
	}

	// 以tcp4协议连接远程服务器，第二个参数本机地址，第三个参数远程地址
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	// n, err := conn.Write([]byte("HEAD / HTTP/1.1\r\n\r\n"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(n)

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // 将 conn 连接中的 数据拷贝到 标准输出 Stdout 当中（当前是命令行控制台）
		log.Println("done")
		done <- struct{}{} // 向主Goroutine发出信号
	}()

	// conn.Close()
	<-done // 等待后台goroutine完成

	// time.Sleep(20 * time.Second)
	// log.Fatal(n)
}
