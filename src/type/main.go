package main

import "fmt"

// NewInt 定义为int类型（实际生产一个新类型，与int类型一致而已）
type NewInt int

// isZero 为 NewInt 定义的“类型方法”(判断是否为0)
func (n NewInt) isZero() bool {
	return n == 0
}

// Add 为 NewInt 定义的“类型方法”
func (n NewInt) Add(m int) int {
	return int(n) + m // 这里 n 的类型和 m 不一致，所以要转换成 int
}

// IntAlias 将int取一个别名叫IntAlias（只是int的别名）
type IntAlias = int

func main() {
	// 将a声明为NewInt类型
	var a NewInt
	// 查看a的类型名
	fmt.Printf("a type: %T\n", a)
	// 将a2声明为IntAlias类型
	var a2 IntAlias
	// 查看a2的类型名
	fmt.Printf("a2 type: %T\n", a2)

	// a type: main.NewInt
	// a2 type: int
	// 结果显示 a 的类型是 main.NewInt，表示 main 包下定义的 NewInt 类型，a2 类型是 int，IntAlias 类型只会在代码中存在，编译完成时，不会有 IntAlias 类型。

	fmt.Println(a.isZero()) // true
	a = 1
	fmt.Println(a.Add(2)) // 3
}
