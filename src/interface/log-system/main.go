package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// 示例，支持多种写入器的日志系统，可以自由扩展多种日志写入设备。
// 这里演示的 是将 日志写入到 文件 和 控制台 当中。

// LogWriter 声明日志写入器接口
type LogWriter interface {
	Write(data interface{}) error
}

// Logger 日志器
type Logger struct {
	// writerList切片 这个日志器用到的日志写入器
	writerList []LogWriter
}

// RegisterWriter 注册一个日志写入器
func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

// Log 将一个data类型的数据写入日志
func (l *Logger) Log(data interface{}) {
	// 遍历所有注册的写入器
	for _, writer := range l.writerList {
		// 将日志输出到每一个写入器中
		writer.Write(data)
	}
}

// NewLogger 创建日志器的实例（方便调用）
func NewLogger() *Logger {
	return &Logger{}
}

func main() {
	// 在程序中使用日志器一般会先通过代码创建日志器（Logger），为日志器添加输出设备（fileWriter、consoleWriter等）。
	// 这些设备中有一部分需要一些参数设定，如文件日志写入器需要提供文件名（fileWriter 的 SetFile() 方法）。
	logger := CreateLogger()
	logger.Log("你好！")
	logger.Log("看看这里！")
}

// CreateLogger 用 fileWriter 和 consoleWriter 创建一个 日志器 Logger
func CreateLogger() *Logger {
	loggers := NewLogger()

	// 文件写入器
	f := newFileWriter()
	// 设置文件名
	if err := f.SetFile("./" + time.Now().Format("2006-01-02") + ".log"); err != nil {
		fmt.Println(err)
	}
	// 将文件写入器注册到 日志器中
	loggers.RegisterWriter(f)

	// 命令行写入器
	c := newConsoleWriter()
	// 将命名行写入器注册到 日志器中
	loggers.RegisterWriter(c)

	return loggers
}

// 1、文件写入器（fileWriter）是众多日志写入器（LogWriter）中的一种。文件写入器的功能是根据一个文件名创建日志文件（fileWriter 的 SetFile 方法）。在有日志写入时，将日志写入文件中。

// fileWriter 声明文件写入器
type fileWriter struct {
	// 在结构体中保存一个文件句柄，以方便每次写入时操作
	file *os.File
}

// SetFile 设置文件写入器写入的文件名
// 文件写入器通过文件名创建文件，这里通过 SetFile 的参数提供一个文件名，并创建文件。
func (f *fileWriter) SetFile(filename string) (err error) {
	// 如果文件已经打开, 关闭前一个文件
	if f.file != nil {
		f.file.Close()
	}
	// 创建一个文件并保存文件句柄
	f.file, err = os.Create(filename)
	// 如果创建的过程出现错误, 则返回错误
	return err
}

// Write 实现 LogWriter接口的 Write()方法
func (f *fileWriter) Write(data interface{}) error {
	// 日志文件可能没有创建成功
	if f.file == nil {
		// 日志文件没有准备好
		return errors.New("file not created")
	}
	// 将数据序列化为字符串
	str := fmt.Sprintf("%v\n", data)
	// 将数据以字节数组写入文件中
	_, err := f.file.Write([]byte(str))
	return err
}

// newFileWriter 创建文件写入器实例（方便调用）
func newFileWriter() *fileWriter {
	return &fileWriter{}
}

// 2、 命令行写入器
// 在 UNIX 的思想中，一切皆文件。文件包括内存、磁盘、网络和命令行等。这种抽象方法方便我们访问这些看不见摸不着的虚拟资源。
// 命令行在 Go语言中也是一种文件，os.Stdout 对应标准输出，一般表示屏幕，也就是命令行，也可以被重定向为打印机或者磁盘文件；
// os.Stderr 对应标准错误输出，一般将错误输出到日志中，不过大多数情况，os.Stdout 会与 os.Stderr 合并输出；
// os.Stdin 对应标准输入，一般表示键盘。
// os.Stdout、os.Stderr、os.Stdin 都是 *os.File 类型，和文件一样实现了 io.Writer 接口的 Write() 方法。

// consoleWriter 命令行写入器
type consoleWriter struct{}

// Write 实现LogWriter的Write()方法
func (f *consoleWriter) Write(data interface{}) error {
	// 将数据序列化为字符串
	str := fmt.Sprintf("%v\n", data)
	// 将数据以字节数组写入命令行中
	_, err := os.Stdout.Write([]byte(str))
	return err
}

// newConsoleWriter 创建命令行写入器实例（方便调用）
func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}
