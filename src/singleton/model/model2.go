package model

// 2) 饿汉式
// 直接创建好对象，不需要判断为空，同时也是线程安全，唯一的缺点是在导入包的同时会创建该对象，并持续占有在内存中。

// Go语言饿汉式可以使用 init 函数，也可以使用全局变量。

// 使用init函数方式：

type tool struct{}

var instance1 *tool

// init 在初始化的时候 实例化 对象
func init() {
	instance1 = new(tool)
}

// GetInstance1 提供获取对象
func GetInstance1() *tool {
	return instance1
}

//全局变量 的方式:
type config struct{}

var cfg *config = new(config)

// NewConfig 提供获取实例的方法
func NewConfig() *config {
	return cfg
}
