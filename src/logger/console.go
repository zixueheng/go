package logger

import (
	"fmt"
	"os"
	"time"
)

// 命令行写入器
// 在 UNIX 的思想中，一切皆文件。文件包括内存、磁盘、网络和命令行等。这种抽象方法方便我们访问这些看不见摸不着的虚拟资源。
// 命令行在 Go语言中也是一种文件，os.Stdout 对应标准输出，一般表示屏幕，也就是命令行，也可以被重定向为打印机或者磁盘文件；
// os.Stderr 对应标准错误输出，一般将错误输出到日志中，不过大多数情况，os.Stdout 会与 os.Stderr 合并输出；
// os.Stdin 对应标准输入，一般表示键盘。
// os.Stdout、os.Stderr、os.Stdin 都是 *os.File 类型，和文件一样实现了 io.Writer 接口的 Write() 方法。

// consoleWriter 命令行写入器
type consoleWriter struct{}

// Write 实现LogWriter的Write()方法
func (f *consoleWriter) Write(upLevel LogLevel, data interface{}) error {
	// 获取初始调用者信息
	fileName, funcName, lineNo := getCaller(4)
	// 将数据序列化为字符串
	str := fmt.Sprintf("[%s] [%s:%s:%d] [%s] %v\n", time.Now().Format("2006-01-02 15:04:05"), fileName, funcName, lineNo, upLevel.ToString(), data)
	// 将数据以字节数组写入命令行中
	_, err := os.Stdout.Write([]byte(str))
	return err
}

// newConsoleWriter 创建命令行写入器实例（方便调用）
func newConsoleWriter() *consoleWriter {
	return &consoleWriter{}
}

// // log 输出日志
// func log(logger *Logger, level LogLevel, msg string) {
// 	// level要大于等于Logger日志器的日志等级，才能进行输出
// 	if level >= logger.level {
// 		// 获取初始调用者信息
// 		fileName, funcName, lineNo := getCaller(3)
// 		fmt.Printf("[%s] [%s:%s:%d] [%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), fileName, funcName, lineNo, level.ToString(), msg)
// 	}
// }

// // Debug 函数
// func (l *Logger) Debug(msg string) {
// 	log(l, Debug, msg)
// }

// // Trace 函数
// func (l *Logger) Trace(msg string) {
// 	log(l, Trace, msg)
// }

// // Info 函数
// func (l *Logger) Info(msg string) {
// 	log(l, Info, msg)
// }

// // Warning 函数
// func (l *Logger) Warning(msg string) {
// 	log(l, Warning, msg)
// }

// // Error 函数
// func (l *Logger) Error(msg string) {
// 	log(l, Error, msg)
// }

// // Fatal 函数
// func (l *Logger) Fatal(msg string) {
// 	log(l, Fatal, msg)
// }
