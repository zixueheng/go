package class1

import (
	"classFactory/base"
)

// Class1 类1
type Class1 struct{}

// 实现Class接口的 String()方法
func (c *Class1) String() string {
	return "Class1"
}

func init() {
	// 启动时 注册 Class1 类的 生成函数
	base.Register("Class1", func() base.Class { // Class1 结构体 实现了 Class 接口，所以返回值类型是 base.Class
		return new(Class1) // 返回 Class1的 实例
	})
}
