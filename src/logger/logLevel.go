package logger

// LogLevel 日志级别
type LogLevel uint8

// 日志级别常量
const (
	DEBUG LogLevel = iota
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

// ToString 将LogLevel转成字符串
func (level LogLevel) ToString() string {
	switch level {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return ""
}
