package logger

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

// LogWriter 声明日志写入器接口
type LogWriter interface {
	// Write 实现写日志的操作
	Write(upLevel LogLevel, data interface{}) error
}

// Logger 日志器
type Logger struct {
	// writerList切片 这个日志器用到的日志写入器
	writerList []LogWriter
	// level 日志等级，日志器的默认等级
	level LogLevel
}

// RegisterWriter 注册一个日志写入器
func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writerList = append(l.writerList, writer)
}

// Log 将一个data类型的数据写入日志，upLevel
func (l *Logger) Log(upLevel LogLevel, data interface{}) {
	// 遍历所有注册的写入器
	for _, writer := range l.writerList {
		// upLevel要大于等于Logger日志器的日志等级，才能进行输出
		if upLevel >= l.level {
			// 将日志输出到每一个写入器中
			writer.Write(upLevel, data)
		}
	}
}

// Debug 函数
func (l *Logger) Debug(msg string) {
	l.Log(DEBUG, msg)
}

// Trace 函数
func (l *Logger) Trace(msg string) {
	l.Log(TRACE, msg)
}

// Info 函数
func (l *Logger) Info(msg string) {
	l.Log(INFO, msg)
}

// Warning 函数
func (l *Logger) Warning(msg string) {
	l.Log(WARNING, msg)
}

// Error 函数
func (l *Logger) Error(msg string) {
	l.Log(ERROR, msg)
}

// Fatal 函数
func (l *Logger) Fatal(msg string) {
	l.Log(FATAL, msg)
}

// NewLogger 是Logger构造函数
func NewLogger(level LogLevel) *Logger {
	return &Logger{
		level: level,
	}
}

// CreateLogger 用 fileWriter 和 consoleWriter 创建一个 日志器 Logger
// level 表示当前日志器的等级
func CreateLogger(level LogLevel) *Logger {
	loggers := NewLogger(level)

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

// getCaller 获取调用信息，返回调用者：文件名、函数名、行号
func getCaller(skip int) (string, string, int) {
	pc, file, line, ok := runtime.Caller(skip) // 2 表示往上2层找调用者
	if !ok {
		return "", "", 0
	}
	funcName := runtime.FuncForPC(pc).Name() //调用者的函数名
	// fmt.Println(funcName, file, line)
	return path.Base(file), funcName, line
}
