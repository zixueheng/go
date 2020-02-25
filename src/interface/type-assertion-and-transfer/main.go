package main

import (
	"fmt"
)

func main() {
	// 类型断言是一个使用在接口值上的操作。语法上它看起来像 i.(T) 被称为断言类型，这里 i 表示一个接口的类型和 T 表示一个类型。
	// 一个类型断言检查它操作对象的动态类型是否和断言的类型匹配。

	// 类型断言的基本格式如下：
	// 	t := i.(T)
	// 其中，i 代表接口变量，T 代表转换的目标类型，t 代表转换后的变量。

	// 这里有两种可能。
	// 第一种，如果断言的类型 T 是一个具体类型，然后类型断言检查 i 的动态类型是否和 T 相同。如果这个检查成功了，类型断言的结果是 i 的动态值，当然它的类型是 T。
	// 换句话说，具体类型的类型断言从它的操作对象中获得具体的值。如果检查失败，接下来这个操作会抛出 panic。例如：
	// var w io.Writer
	// w = os.Stdout
	// f := w.(*os.File) // 成功: f == os.Stdout
	// fmt.Printf("%v", *f)
	// c := w.(*bytes.Buffer) // 死机：接口保存*os.file，而不是*bytes.buffer
	// fmt.Printf("%v", *c)

	// 第二种，如果相反断言的类型 T 是一个接口类型，然后类型断言检查是否 i 的动态类型满足 T。如果这个检查成功了，动态值没有获取到；这个结果仍然是一个有相同类型和值部分的接口值，但是结果有类型 T。
	// 换句话说，对一个接口类型的类型断言改变了类型的表述方式，改变了可以获取的方法集合（通常更大），但是它保护了接口值内部的动态类型和值的部分。
	// 在下面的第一个类型断言后，w 和 rw 都持有 os.Stdout 因此它们每个有一个动态类型 *os.File，但是变量 w 是一个 io.Writer 类型只对外公开出文件的 Write 方法，然而 rw 变量也只公开它的 Read 方法。
	// var w io.Writer
	// w = os.Stdout
	// rw := w.(io.ReadWriter) // 成功：*os.file具有读写功能
	// w = new(ByteCounter)
	// rw = w.(io.ReadWriter) // 死机：*字节计数器没有读取方法

	// 将接口转换为其他接口

	// 创建动物的名字到实例的映射
	// 将鸟和猪的实例创建后，被保存到 interface{} 类型的 map 中。interface{} 类型表示空接口，意思就是这种接口可以保存为任意类型。
	animals := map[string]interface{}{
		"bird": new(bird),
		"pig":  new(pig),
	}
	// 遍历映射
	for name, obj := range animals {
		// 对保存有鸟或猪的实例的 interface{} 变量进行断言操作，如果断言对象是断言指定的类型，则返回转换为断言对象类型的接口；
		// 如果不是指定的断言类型时，断言的第二个参数将返回 false

		// 判断对象是否为飞行动物，断言成功 f 就是 Flyer 类型
		f, isFlyer := obj.(Flyer)
		// 判断对象是否为行走动物，断言成功 w 就是 Walker 类型
		w, isWalker := obj.(Walker)
		fmt.Printf("name: %s isFlyer: %v isWalker: %v\n", name, isFlyer, isWalker)

		// 如果是飞行动物则调用飞行动物接口
		if isFlyer {
			f.Fly()
		}
		// 如果是行走动物则调用行走动物接口
		if isWalker {
			w.Walk()
		}
	}
	// name: pig isFlyer: false isWalker: true
	// pig: walk
	// name: bird isFlyer: true isWalker: true
	// bird: fly
	// bird: walk

	// 将接口转换为其他类型
	// 可以实现将接口转换为普通的指针类型。例如将 Walker 接口转换为 *pig 类型：
	p1 := new(pig)
	var a Walker = p1                 // 由于 pig 实现了 Walker 接口，因此可以被隐式转换为 Walker 接口类型保存于 a 中
	p2 := a.(*pig)                    // 这里实现接口 a 转换成 指针类型 *pig，由于 a 中保存的本来就是 *pig 本体，因此可以转换为 *pig 类型
	fmt.Printf("p1=%p p2=%p", p1, p2) // 打印 p1 和 p2 指针是相同的
	// p1=0x597c18 p2=0x597c18

}

// Flyer 定义飞行动物接口
type Flyer interface {
	Fly()
}

// Walker 定义行走动物接口
type Walker interface {
	Walk()
}

// bird 定义鸟类
type bird struct{}

// Fly() 实现飞行动物接口
func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

// Walk() 实现行走动物接口
func (b *bird) Walk() {
	fmt.Println("bird: walk")
}

// pig 定义猪
type pig struct{}

// Walk(), 实现行走动物接口
func (p *pig) Walk() {
	fmt.Println("pig: walk")
}
