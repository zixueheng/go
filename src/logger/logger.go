package logger

import (
	"path"
	"runtime"
)

// Logger 结构体
type Logger struct {
	level LogLevel
}

// LogLevel 日志级别
type LogLevel uint8

// 日志级别常量
const (
	Debug LogLevel = iota
	Trace
	Info
	Warning
	Error
	Fatal
)

// NewLogger 是Logger构造函数
func NewLogger(level LogLevel) *Logger {
	return &Logger{
		level,
	}
}

// ToString 将LogLevel转成字符串
func (level LogLevel) ToString() string {
	switch level {
	case Debug:
		return "Debug"
	case Trace:
		return "Trace"
	case Info:
		return "Info"
	case Warning:
		return "Warning"
	case Error:
		return "Error"
	case Fatal:
		return "Fatal"
	}
	return ""
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
