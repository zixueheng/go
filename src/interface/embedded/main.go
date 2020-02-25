package main

import "io"

// 一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。
// 只要接口的所有方法被实现，则这个接口中的所有嵌套接口的方法均可以被调用。

// Go语言的 io 包中定义了写入器（Writer）、关闭器（Closer）和写入关闭器（WriteCloser）3个接口

// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }
// type Closer interface {
// 	Close() error
// }
// 由 Writer 和 Closer 两个接口嵌入。也就是说，WriteCloser 同时拥有了 Writer 和 Closer 的特性
// type WriteCloser interface {
// 	Writer
// 	Closer
// }

// device 声明一个设备结构
type device struct{}

// 实现io.Writer的Write()方法
func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}

// 实现io.Closer的Close()方法
func (d *device) Close() error {
	return nil
}

func main() {
	// 由于 device 实现了 io.WriteCloser 的所有嵌入接口，因此 device 指针就会被隐式转换为 io.WriteCloser 接口
	var wc io.WriteCloser = new(device)

	// 调用了 wc（io.WriteCloser接口）的 Write() 方法，由于 wc 被赋值 *device，因此最终会调用 device 的 Write() 方法。
	wc.Write(nil)
	// 调用了 wc（io.WriteCloser接口）的 Close() 方法，由于 wc 被赋值 *device，因此最终会调用 device 的 Close() 方法。
	wc.Close()

	var writeOnly io.Writer = new(device)
	writeOnly.Write(nil)
	// writeOnly.Close() // 这里由于 Writer 没有 Close() 方法 而报错

}
