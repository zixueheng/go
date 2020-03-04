package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 创建一个程序结束码的通道
	exitChan := make(chan int)
	// 将服务器并发运行
	go server("127.0.0.1:7001", exitChan)
	// 通道阻塞, 等待接收返回值
	code := <-exitChan
	// 标记程序返回值并退出
	os.Exit(code)
}

// 编译所有代码并运行，命令行提示如下：
// listen: 127.0.0.1:7001
// 此时，Socket 侦听成功。在操作系统中的命令行中输入：
// telnet 127.0.0.1 7001
// 在 Telnet 连接后，输入字符串 hello 查看效果
// 当输入 @close 时 关闭会话
// 当输入 @shutdown 时，服务器关闭
// 注意：这里可以开两个 CMD 窗口测试效果，两个运行互不干扰，在一个里面输入@close关闭本次会话另外一个还在，输入@shutdown 则服务器关闭所有会员都会关闭

// 服务逻辑, 传入地址和退出的通道
func server(address string, exitChan chan int) {
	// 根据给定地址进行侦听
	l, err := net.Listen("tcp", address)
	// 如果侦听发生错误, 打印错误并退出
	if err != nil {
		fmt.Println(err.Error())
		exitChan <- 1
	}
	// 打印侦听地址, 表示侦听成功
	fmt.Println("listen: " + address)
	// 延迟关闭侦听器
	defer l.Close()
	// 侦听循环
	for {
		// 服务器接受了一个连接。在没有连接时，Accept() 函数调用后会一直阻塞。连接到来时，返回 conn 和错误变量，conn 的类型是 *tcp.Conn
		conn, err := l.Accept()
		// 某些情况下，连接接受会发生错误，不影响服务器逻辑，这时重新进行新连接接受
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		// 根据连接开启会话, 这个过程需要并行执行
		// 每个连接会生成一个会话。这个会话的处理与接受逻辑需要并行执行，彼此不干扰
		go handleSession(conn, exitChan)
	}
}

// 连接的会话逻辑
// 回音服务器的基本逻辑是“收到什么返回什么”，reader.ReadString 可以一直读取 Socket 连接中的数据直到碰到期望的结尾符。这种期望的结尾符也叫定界符，一般用于将 TCP 封包中的逻辑数据拆分开。
// 下例中使用的定界符是回车换行符（“\r\n”），HTTP 协议也是使用同样的定界符。使用 reader.ReadString() 函数可以将封包简单地拆分开。
func handleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("Session started:")
	// 创建一个网络连接数据的读取器
	reader := bufio.NewReader(conn)
	// 接收数据的循环
	for {
		// 读取字符串, 直到碰到回车返回
		str, err := reader.ReadString('\n') // 当没有数据时，调用 reader.ReadString 会发生阻塞，等待数据的到来。一旦数据到来，就可以进行各种逻辑处理
		// 数据读取正确
		if err == nil {
			// 去掉字符串尾部的回车
			str = strings.TrimSpace(str) // reader.ReadString 读取返回的字符串尾部带有回车符，使用 strings.TrimSpace() 函数将尾部带的回车和空白符去掉
			// 处理Telnet指令
			if !processTelnetCommand(str, exitChan) {
				conn.Close()
				break
			}
			// Echo逻辑, 发什么数据, 原样返回
			conn.Write([]byte(str + "\r\n")) // 将有效数据通过 conn 的 Write() 方法写入，同时在字符串尾部添加回车换行符（“\r\n”），数据将被 Socket 发送给连接方
		} else {
			// 发生错误
			fmt.Println("Session closed")
			conn.Close()
			break
		}
	}
}

// Telnet命令处理
// 输入“@close”退出当前连接会话。
// 输入“@shutdown”终止服务器运行。
func processTelnetCommand(str string, exitChan chan int) bool {
	// @close指令表示终止本次会话
	if strings.HasPrefix(str, "@close") {
		fmt.Println("Session closed")
		// 告诉外部需要断开连接
		return false
		// @shutdown指令表示终止服务进程
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("Server shutdown")
		// 往通道中写入0, 阻塞等待接收方处理
		exitChan <- 0
		// 告诉外部需要断开连接
		return false
	}
	// 打印输入的字符串
	fmt.Println(str)
	return true
}
