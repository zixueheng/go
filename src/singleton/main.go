package main

import (
	"fmt"
	"singleton/model"
)

// 单例模式也叫单子模式，是常用的模式之一，在它的核心结构中只包含一个被称为单例的特殊类，能够保证系统运行中一个类只创建一个实例。

func main() {
	tool := model.GetInstance()
	fmt.Println(tool.String())
}
