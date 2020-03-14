package logger

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// 文件写入器（fileWriter）是众多日志写入器（LogWriter）中的一种。文件写入器的功能是根据一个文件名创建日志文件（fileWriter 的 SetFile 方法）。在有日志写入时，将日志写入文件中。

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
	f.file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	// 如果创建的过程出现错误, 则返回错误
	return err
}

// Write 实现 LogWriter接口的 Write()方法
func (f *fileWriter) Write(upLevel LogLevel, data interface{}) error {
	// 日志文件可能没有创建成功
	if f.file == nil {
		// 日志文件没有准备好
		return errors.New("file not created")
	}
	// 获取初始调用者信息
	fileName, funcName, lineNo := getCaller(4)
	// 将数据序列化为字符串
	str := fmt.Sprintf("[%s] [%s:%s:%d] [%s] %v\n", time.Now().Format("2006-01-02 15:04:05"), fileName, funcName, lineNo, upLevel.ToString(), data)
	// 将数据以字节数组写入文件中
	_, err := f.file.Write([]byte(str))
	return err
}

// newFileWriter 创建文件写入器实例（方便调用）
func newFileWriter() *fileWriter {
	return &fileWriter{}
}
