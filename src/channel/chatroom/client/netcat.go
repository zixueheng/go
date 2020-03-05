// netcat 是一个简单的TCP服务器读/写客户端
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		// 将 conn 连接中的 数据拷贝到 标准输出 Stdout 当中（当前是命令行控制台）
		io.Copy(os.Stdout, conn) // 注意：忽略错误，程序会一致停留在这里，具体什么原因还没理解
		log.Println("done")
		done <- struct{}{} // 向主Goroutine发出信号
	}()

	mustCopy(conn, os.Stdin) // 将 当前命令行中的输入 拷贝到 当前连接conn 中

	conn.Close()
	<-done // 等待后台goroutine完成
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
