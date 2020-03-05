package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// client 类型定义为 单向发送通道（通道类型string）
type client chan<- string

var (
	entering = make(chan client) // 发送 客户端进入提示的通道
	leaving  = make(chan client) // 发送 客户端离开提示的通道

	messages = make(chan string) // 客户端消息通道
)

// 广播消息（包括客户端进入、离开和发送的文本消息）
// 广播的意思是 有消息要发送给所有的客户端
func broadcaster() {
	clients := make(map[client]bool) // make 一个 客户端的 map
	// 无限循环等待消息
	for {
		// 用 select 多路复用处理消息
		select {
		case cli := <-entering: // 每当有消息从 entering 里面发送过来，就生成一个新的 key - value，相当于给 clients 里面增加一个新的 client
			clients[cli] = true
		case cli := <-leaving: // 每当有消息从 leaving 里面发送过来，就删掉这个 key - value 对，并关闭对应的 channel
			delete(clients, cli)
			close(cli)
		case msg := <-messages: // 每当有广播消息从 messages 发送进来，都会循环 cliens 对里面的每个 channel 发消息
			for cli := range clients { // 遍历 map 只有一个 参数"cli" 表示只要 键，不要值(bool)
				// 这里的 cli 实际上就是 handelConn 第74行添加进来的通道(函数内变量 ch)
				// msg 消息发送给 这个ch，有消息进入 ch 之后，在 clientWriter 线程中的 循环接收 并写入当前的 连接 conn 对象
				// 更进一步解释：conn连接对象里面有了新消息，客户端(netcat.go) 会不断拷贝 conn 中的 新消息显示到 客户端的命令行输出中
				cli <- msg
			}
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster() // 开启 广播 线程，它在不停的等待消息 并处理消息

	// 一直循环等待客户端连接，main函数不会结束
	for {
		conn, err := listener.Accept()
		if err != nil { // 发生连接错误打印出来 并 继续
			log.Print(err)
			continue
		}
		go handleConn(conn) // 开启处理 客户端连接的 线程
	}
}

// handleConn 函数会为每个过来处理的 conn 都创建一个新的 channel，开启一个新的 goroutine 去把发送给这个 channel 的消息写进 conn
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch) // ch 是双向通道可以传递给 单向通道（这里参数约定是 接收通道）

	who := conn.RemoteAddr().String()
	ch <- "欢迎 " + who // 生成欢迎消息放入 ch 通道，执行效果：只在自己的窗口里面显示

	messages <- who + " 上线" // 生成一条广播消息写进 messages 里，执行效果：发送给其他所有的客户端连接

	// 注意：这里能把 ch 发送给 entering，是因为 entering 是一个通道，这个通道 可放置 单向发送通道（ch是双向通道，我理解可向下转为单向通道）
	// 把ch 发送给 entering，实际上表示 把这个通道作为 当前客户端的标识
	entering <- ch // 把这个 channel 加入到客户端集合，也就是 entering <- ch

	// 监听客户端往 conn 里写的数据，每扫描到一条就将这条消息发送到广播 channel 中
	input := bufio.NewScanner(conn)
	for input.Scan() { // 按行读取
		messages <- who + ": " + input.Text()
	}
	// 注意：忽略 input.Err() 中可能的错误

	leaving <- ch           // 如果关闭了客户端，那么把队列离开写入 leaving 交给广播函数去删除这个客户端并关闭这个客户端
	messages <- who + " 下线" // 广播通知其他客户端该客户端已关闭
	conn.Close()            // 关闭这个客户端的连接 Conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	// 循环接收 通道里面的数据并写入连接 conn 中
	for msg := range ch {
		fmt.Fprintln(conn, msg) // 注意：忽略网络层面的错误
	}
}
