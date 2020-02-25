package class2

import "classFactory/base"

// Class2 类2
type Class2 struct{}

// 实现Class接口的 String()方法
func (c *Class2) String() string {
	return "Class2"
}

func init() {
	// 启动时 注册 Class2 类的 生成函数
	base.Register("Class2", func() base.Class { // Class2 结构体 实现了 Class 接口，所以返回值类型是 base.Class
		return new(Class2) // 返回 Class2 的实例
	})
}
