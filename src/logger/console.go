package logger

import (
	"fmt"
	"time"
)

// log 输出日志
func log(logger *Logger, level LogLevel, msg string) {
	// level要大于等于Logger日志器的日志等级，才能进行输出
	if level >= logger.level {
		// 获取初始调用者信息
		fileName, funcName, lineNo := getCaller(3)
		fmt.Printf("[%s] [%s:%s:%d] [%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), fileName, funcName, lineNo, level.ToString(), msg)
	}
}

// Debug 函数
func (l *Logger) Debug(msg string) {
	log(l, Debug, msg)
}

// Trace 函数
func (l *Logger) Trace(msg string) {
	log(l, Trace, msg)
}

// Info 函数
func (l *Logger) Info(msg string) {
	log(l, Info, msg)
}

// Warning 函数
func (l *Logger) Warning(msg string) {
	log(l, Warning, msg)
}

// Error 函数
func (l *Logger) Error(msg string) {
	log(l, Error, msg)
}

// Fatal 函数
func (l *Logger) Fatal(msg string) {
	log(l, Fatal, msg)
}
