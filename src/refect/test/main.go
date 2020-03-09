package main

import "fmt"

func main() {
	var fn = func() {
		fmt.Println("hello")
	}

	// 不可以 直接用 字符串 方式调用函数
	// fnName := "fn"
	// fnName() // cannot call non-function fnName (type string)

	// 但是可以通过map 做 函数名字符串 到函数的 映射
	funcs := map[string]func(){
		"fnName": fn,
	}
	funcs["fnName"]() // 可以正常输出 hello

	// 更进一步 将 map 的 值设为空接口试试
	// funcs2 := map[string]interface{}{
	// 	"fnName": fn,
	// }
	// funcs2["fnName"]是接口类型，不能调用
	// funcs2["fnName"]() // cannot call non-function funcs2["fnName"] (type interface {})
	// 所以还是用 reflect 包吧
}
