package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// GetTCPAddr 演示 获取一个 TCPAddr
func GetTCPAddr() {
	// net 参数是 "tcp4"、"tcp6"、"tcp" 中的任意一个，分别表示 TCP(IPv4-only)，TCP(IPv6-only) 或者 TCP(IPv4,IPv6) 中的任意一个；
	// addr 表示域名或者 IP 地址，例如 "c.biancheng.net:80" 或者 "127.0.0.1:22"。
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "localhost:8080")
	fmt.Println(tcpAddr) // 127.0.0.1:8080
}

// 每隔5秒向tcp客户端发送时间
func echo(conn *net.TCPConn) {
	// 定义一个5秒的打点器
	ticker := time.Tick(5 * time.Second)
	// 遍历打点器，每隔5秒返回当前时间
	for now := range ticker {
		// 将当前时间写入当前的TCP连接
		n, err := conn.Write(([]byte)(now.String() + "\r\n")) // 加入换行
		if err != nil {                                       // 发生错误打印错误并关闭连接
			log.Println(err)
			conn.Close()
			return
		}
		// 在当前终端发生发送的信息
		fmt.Printf("send %d bytes to %s\n", n, conn.RemoteAddr())
	}
}

// 使用TCP连接定时向客服端发送当前时间
func demo1() {
	// 创建一个 tcp 地址
	var address net.TCPAddr = net.TCPAddr{
		IP:   net.ParseIP("localhost"), // ParseIP 将字符串转成 net.IP
		Port: 8000,
	}

	listener, err := net.ListenTCP("tcp4", &address) // tcp4协议监听上面的地址
	if err != nil {
		log.Fatal(err)
	}

	// 循环等待客户端连接
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("远程地址：", conn.RemoteAddr())
		go echo(conn)
	}
}

func main() {
	demo1()
	// 运行后，在另外一个终端中 telnet 127.0.0.1 8000，也可以看到每隔5秒显示服务器发过来的时间
	// 另外可以 运行 tcpClient中的 main.go 以程序的方式连接服务器
}
