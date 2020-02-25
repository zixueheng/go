package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// 正则表达式是一种进行模式匹配和文本操纵的复杂而又强大的工具。
// 虽然正则表达式比纯粹的文本匹配效率低，但是它却更灵活，按照它的语法规则，根据需求构造出的正则表达式能够从原始文本中筛选出几乎任何你想要得到的字符组合。

// 匹配指定类型的字符串
func demo1() {
	buf := "abc azc a7c aac 888 a9c  tac"
	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(`a\dc`) //构造 a数字c 的解析器
	if reg1 == nil {
		fmt.Println("解析器构建失败")
		return
	}
	strs := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println(strs)
	// [[a7c] [a9c]]
	for _, v := range strs {
		fmt.Printf("Value: %v, Type: %T\n", v, v)
	}
	// Value: [a7c], Type: []string
	// Value: [a9c], Type: []string
}

// 匹配 a 和 c 中间包含一个数字的字符串
func demo2() {
	buf := "abc azc a7c aac 888 a9c  tac"
	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(`a[0-9]c`)
	if reg1 == nil { //解释失败，返回nil
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("result1 = ", result1)
}

// 过 Compile 方法返回一个 Regexp 对象，实现匹配，查找，替换相关的功能
func demo3() {
	//目标字符串
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+" //正则

	// f 一个转换函数
	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}

	re, _ := regexp.Compile(pat)

	// 将匹配到的部分替换为 "##.#"
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str) // John: ##.# William: ##.# Steve: ##.#

	// 将匹配的字符串用 函数进行转换
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2) // John: 5156.68 William: 9134.46 Steve: 11264.36
}

func main() {
	// demo1()
	demo3()
}
