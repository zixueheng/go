package main

import (
	"fmt"
	"io"
	"time"
)

// 1、一个类型可以实现多个接口
// 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。

// Socket 套接字类型
type Socket struct{}

// Writer 接口，io包中的
// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

// Closer 接口，io包中的
// type Closer interface {
// 	Close() error
// }

// Write 方法实现了 io.Writer 接口
func (s *Socket) Write(p []byte) (n int, err error) {
	fmt.Println("写入数据", len(p))
	return 0, nil
}

// Close 方法实现了 io.Closer 接口
func (s *Socket) Close() error {
	fmt.Println("关闭通道")
	return nil
}

// 使用 Socket 实现的 Writer 接口的代码（22行），无须了解 Writer 接口的实现者是否具备 Closer 接口的特性。
// 同样，使用 Closer 接口的代码也并不知道 Socket 已经实现了 Writer 接口

// 使用io.Writer的代码, 并不知道Socket和io.Closer的存在
func usingWriter(writer io.Writer) {
	// str := fmt.Sprintf("%v", "测试数据")
	writer.Write([]byte("测试数据"))
}

// 使用io.Closer, 并不知道Socket和io.Writer的存在
func usingCloser(closer io.Closer) {
	closer.Close()
}

func main() {
	// 实例化Socket
	s := new(Socket)
	// Socket 类型实现了 io.Writer 和 io.Closer 接口中定义的方法，所以下面的代码是可行的
	usingWriter(s)
	usingCloser(s)
	// 写入数据 12
	// 关闭通道

	g := new(GameService)
	// g 就可以使用 Start() 方法和 Log() 方法，其中，Start() 由 GameService 实现，Log() 方法由 Logger 实现。
	g.Start()
	g.Log("开始游戏于：" + time.Now().Format("2006-01-02 15:04:05"))
	// 开始游戏
	// 记录日志： 开始游戏于：2020-01-06 14:44:37
}

// 2、多个类型可以实现相同的接口
// 一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。
// 也就是说，使用者并不关心某个接口的方法是通过一个类型完全实现的，还是通过多个结构嵌入到一个结构体中拼凑起来共同实现的。

// Service 一个服务需要满足能够开启和写日志的功能
type Service interface {
	Start()     // 开启服务
	Log(string) // 日志输出
}

// Logger 日志器
type Logger struct{}

// Log 实现Service的Log()方法
func (g *Logger) Log(l string) {
	fmt.Println("记录日志：", l)
}

// GameService 游戏服务
type GameService struct {
	Logger // 嵌入日志器
}

// Start 实现Service的Start()方法
func (g *GameService) Start() {
	fmt.Println("开始游戏")
}

// 代码说明：
// Service 接口定义了两个方法：一个是开启服务的方法（Start()），一个是输出日志的方法（Log()）。
// 使用 GameService 结构体来实现 Service，GameService 自己的结构只能实现 Start() 方法，
// 而 Service 接口中的 Log() 方法已经被一个能输出日志的日志器（Logger）实现了，无须再进行 GameService 封装，或者重新实现一遍。
// 所以，选择将 Logger 嵌入到 GameService 能最大程度地避免代码冗余，简化代码结构。
