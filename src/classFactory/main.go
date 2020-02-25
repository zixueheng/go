package main

import (
	"classFactory/base"
	_ "classFactory/class1"
	_ "classFactory/class2"
	"fmt"
)

// Go 的工厂类 示例
// 通过匿名引用方法导入了 class1 和 class2 两个包。在 main() 函数调用前，这两个包的 init() 函数会被自动调用，从而自动注册 Class1 和 Class2。
// 然后我们就可以在 main函数中 通过类名的字符串创建 对应类的 实例 了

func main() {
	class1 := base.Create("Class1")           // 这里传的字符串 Class1 是 class1 包中 init()函数注册的类
	fmt.Printf("Name: %s\n", class1.String()) // Name: Class1

	class2 := base.Create("Class2")
	fmt.Printf("Name: %s\n", class2.String()) // Name: Class2
}
