package main

import (
	"fmt"
	"strings"
)

// 对数据的操作进行多步骤的处理被称为链式处理，本例中使用多个字符串作为数据集合，然后对每个字符串进行一系列的处理，
// 用户可以通过系统函数或者自定义函数对链式处理中的每个环节进行自定义。

// StringProcess 字符串处理函数，传入字符串切片和处理函数切片
func StringProcess(stringList []string, funcChain []func(string) string) {
	for index, str := range stringList {
		result := str
		for _, funcName := range funcChain {
			result = funcName(result)
		}
		stringList[index] = result
	}
}

// StringTrimPrefix 自定义去除字符串前缀函数
func StringTrimPrefix(str string) string {
	return strings.TrimPrefix(str, "go")
}

func main() {
	// 待处理的字符串列表
	list := []string{
		"go scanner ",
		" go parser",
		" go compiler",
		"go printer ",
		"  go formater  ",
	}

	funcChain := []func(string) string{
		strings.TrimSpace, // 去除字符串两边空格
		StringTrimPrefix,  // 去除 前缀 go
		strings.TrimSpace, // 再去除两边空格
		strings.ToUpper,   // 转成大写
	}

	StringProcess(list, funcChain)

	// 输出处理好的字符串
	for _, str := range list {
		fmt.Println(str)
	}
	// SCANNER
	// PARSER
	// COMPILER
	// PRINTER
	// FORMATER
}
