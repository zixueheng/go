package main

import "fmt"

// 每个接口类型由数个方法组成。接口的形式代码如下：
//     type 接口类型名 interface{
//         方法名1( 参数列表1 ) 返回值列表1
//         方法名2( 参数列表2 ) 返回值列表2
//         …
//     }
// 接口类型名：使用 type 将接口定义为自定义的类型名。Go语言的接口在命名时，一般会在单词后面添加 er，如有写操作的接口叫 Writer，有字符串功能的接口叫 Stringer，有关闭功能的接口叫 Closer 等。
// 方法名：当方法名首字母是大写时，且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
// 参数列表、返回值列表：参数列表和返回值列表中的参数变量名可以被忽略

// DataWriter 接口
type DataWriter interface {
	// WriteData 方法 输入一个 interface{} 类型的 data，返回一个 error 结构表示可能发生的错误。
	WriteData(data interface{}) error

	CanWrite() bool
}

// file 结构体即文件结构，用于实现DataWriter
type file struct{}

// 实现 DataWriter 接口的 WriteData 方法（file结构体的方法）
func (f *file) WriteData(data interface{}) error {
	fmt.Println("Data:", data)
	return nil
}

// 实现 DataWriter 接口的 CanWrite 方法（file结构体的方法）
// 如果没有实现 DataWriter 接口 的 CanWrite 方法，编译不会通过
func (f *file) CanWrite() bool {
	return true
}

// 接口被实现的条件：
// 一、接口的方法与实现接口的类型方法格式一致
// 二、接口中所有方法均被实现

func main() {
	// 实例化一个 file
	f := new(file)
	// 声明一个 DataWriter 类型的接口
	var writer DataWriter

	// 将接口赋值f，也就是*file类型
	// 关键的地方：将 *file 类型的 f 赋值给 DataWriter 接口的 writer，虽然两个变量类型不一致。但是 writer 是一个接口，且 f 已经完全实现了 DataWriter() 的所有方法，因此赋值是成功的。
	writer = f

	// 使用DataWriter接口进行数据写入
	writer.WriteData("something")
	// Data: something

}
