package main

import "fmt"

// 切片（slice）是对数组的一个连续片段的引用，所以切片是一个引用类型（因此更类似于 C/C++ 中的数组类型，或者 Python 中的 list 类型），
// 这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集，需要注意的是，终止索引标识的项不包括在切片内。

func main() {
	// 1、从数组或切片生成新的切片
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])
	// [1 2 3] [2]
	// 语法说明如下：
	// slice：表示目标切片对象；
	// 开始位置：对应目标切片对象的索引；
	// 结束位置：对应目标切片的结束索引。

	// 从数组或切片生成新的切片拥有如下特性：
	// 取出的元素数量为：结束位置 - 开始位置；
	// 取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取；
	// 当缺省开始位置时，表示从连续区域开头到结束位置；
	// 当缺省结束位置时，表示从开始位置到整个连续区域末尾；
	// 两者同时缺省时，与切片本身等效；
	// 两者同时为 0 时，等效于空切片，一般用于切片复位。

	// 2、直接声明新的切片
	// 除了可以从原有的数组或者切片中生成切片外，也可以声明一个新的切片，每一种类型都可以拥有其切片类型，表示多个相同类型元素的连续集合，因此切片类型也可以被声明：
	// var name []Type
	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	// 声明一个空切片，这里已分配内存，只是还没有元素
	var numListEmpty = []int{}
	// 输出3个切片
	fmt.Println(strList, numList, numListEmpty)
	// [] [] []
	// 输出3个切片大小
	fmt.Println(len(strList), len(numList), len(numListEmpty))
	// 0 0 0
	// 切片判定空的结果
	fmt.Println(strList == nil)
	fmt.Println(numList == nil)
	fmt.Println(numListEmpty == nil)
	// true
	// true
	// false

	// 3、使用 make() 函数构造切片
	// make( []Type, size, cap )
	// 其中 Type 是指切片的元素类型，size 指的是为这个类型分配多少个元素，cap 为预分配的元素数量，这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题。
	x := make([]int, 2)
	y := make([]int, 2, 10)
	fmt.Println(x, y)
	fmt.Println(len(x), len(y))
	// [0 0] [0 0]
	// 2 2
	// 注意：
	// 使用 make() 函数生成的切片一定发生了内存分配操作，
	// 但给定开始与结束位置（包括切片复位）的切片只是将新的切片结构指向已经分配好的内存区域，设定开始与结束位置，不会发生内存分配操作。

	// append() 可以为切片动态添加元素
	var m []int
	fmt.Println(m)
	// []
	m = append(m, 1) // 追加1个元素
	fmt.Println(m)
	// [1]
	m = append(m, 1, 2, 3) // 追加多个元素, 手写解包方式
	fmt.Println(m)
	// [1 1 2 3]
	m = append(m, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包
	fmt.Println(m)
	// [1 1 2 3 1 2 3]

	// copy()：切片复制（切片拷贝）
	// copy() 可以将一个数组切片复制到另一个数组切片中，如果加入的两个数组切片不一样大，就会按照其中较小的那个数组切片的元素个数进行复制。
	// 格式：copy( destSlice, srcSlice []T) int
	// 其中 srcSlice 为数据来源切片，destSlice 为复制的目标（也就是将 srcSlice 复制到 destSlice），
	// 目标切片必须分配过空间且足够承载复制的元素个数，并且来源和目标的类型必须一致，copy() 函数的返回值表示实际发生复制的元素个数。
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8}
	n1 := copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中，因为两个元素最小的个数是3
	fmt.Println(slice1, slice2, n1)
	// [1 2 3 4 5] [1 2 3] 3
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := []int{6, 7, 8}
	n2 := copy(slice3, slice4) // 只会复制slice4的3个元素到slice2的前3个位置，因为两个元素最小的个数是3
	fmt.Println(slice3, slice4, n2)
	// [6 7 8 4 5] [6 7 8] 3

	// 创建一个整型切片，并赋值
	slice5 := []int{10, 20, 30, 40}
	// 迭代每个元素，并显示值和地址，
	// range 返回的是每个元素的副本，而不是直接返回对该元素的引用
	for index, value := range slice5 {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n", value, &value, &slice5[index])
	}
	// Value: 10 Value-Addr: C000010260 ElemAddr: C00000E500
	// Value: 20 Value-Addr: C000010260 ElemAddr: C00000E508
	// Value: 30 Value-Addr: C000010260 ElemAddr: C00000E510
	// Value: 40 Value-Addr: C000010260 ElemAddr: C00000E518
}
