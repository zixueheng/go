package main

import (
	"logger"
)

// 日志需求：
// 1）支持往不同位置输出
// 2）日志需要分级别（Debug Trace Info Warning Error Fatal）
// 3）日志要支持开关（什么级别的日志才可以记录，如只记录Info以上级别的日志）
// 4）日志格式要包括：时间、级别、文件|行号、日志内容
// 5）日志要按日期文件大小切割

func main() {
	// 使用系统内置的日志库进行日志打印
	// 打印到终端中
	// log.SetOutput(os.Stdout) // log 默认是标准输出(即: 终端)
	// log.Println("日志记录") // 2020/02/12 17:46:04 日志记录
	// file, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	// if err != nil {
	// 	fmt.Printf("Error: %v", err)
	// }
	// log.SetOutput(file)      // 设置输入到文件中
	// log.Println("记录到文件中的日志") // 这里会输出到 log.log 文件中

	// 使用自己编写的logger包
	logger := logger.CreateLogger(logger.WARNING)
	logger.Debug("一条Debug日志")
	logger.Trace("一条Trace日志")
	logger.Info("一条Info日志")
	logger.Warning("一条Warning日志2")
	logger.Error("一条Error日志2")
	logger.Fatal("一条Fatal日志2")
}
