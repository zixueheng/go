package main

import (
	"fmt"
	"sort"
)

// Go语言的 sort.Sort 函数使用了一个接口类型 sort.Interface 来指定通用的排序算法和可能被排序到的序列类型之间的约定。
// 这个接口的实现由序列的具体表示和它希望排序的元素决定，序列的表示经常是一个切片。
// 一个内置的排序算法需要知道三个东西：序列的长度，表示两个元素比较的结果，一种交换两个元素的方式；这就是 sort.Interface 的三个方法：
//     package sort
//     type Interface interface {
//         Len() int            // 获取元素数量
//         Less(i, j int) bool // i，j是序列元素的指数。
//         Swap(i, j int)        // 交换元素
//     }

// 使用sort.Interface接口进行排序
// 对一系列字符串进行排序时，使用字符串切片（[]string）承载多个字符串。使用 type 关键字，将字符串切片（[]string）定义为自定义类型 MyStringList。
// 为了让 sort 包能识别 MyStringList，能够对 MyStringList 进行排序，就必须让 MyStringList 实现 sort.Interface 接口：

// MyStringList 定义为 字符串切片类型
// 接口实现不受限于结构体，任何类型都可以实现接口。要排序的字符串切片 []string 是系统定制好的类型，无法让这个类型去实现 sort.Interface 排序接口。因此，需要将 []string 定义为自定义的类型。
type MyStringList []string

// Len 获取元素数量
func (m MyStringList) Len() int {
	return len(m)
}

// Less 比较元素
func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

// Swap 交换元素
func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	strlist := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	// 由于将 []string 定义成 MyStringList 类型，字符串切片初始化的过程等效于下面的写法：
	// strlist := []string {
	//     "3. Triple Kill",
	//     "5. Penta Kill",
	//     "2. Double Kill",
	//     "4. Quadra Kill",
	//     "1. First Blood",
	// }

	// 使用 sort 包的 Sort() 函数，将 strlist（MyStringList类型）进行排序。
	// 排序时，sort 包会通过 MyStringList 实现的 Len()、Less()、Swap() 这 3 个方法进行数据获取和修改
	sort.Sort(strlist)

	for _, v := range strlist {
		fmt.Println(v)
	}
	// 1. First Blood
	// 2. Double Kill
	// 3. Triple Kill
	// 4. Quadra Kill
	// 5. Penta Kill

	// 常见类型的便捷排序：
	// 1) 字符串切片的便捷排序
	// sort 包中有一个 StringSlice 类型，定义如下：
	// 	type StringSlice []string
	// 	func (p StringSlice) Len() int           { return len(p) }
	// 	func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
	// 	func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
	// 	// Sort is a convenience method.
	// 	func (p StringSlice) Sort() { Sort(p) }

	// sort 包中的 StringSlice 的代码与 MyStringList 的实现代码几乎一样。因此，只需要使用 sort 包的 StringSlice 就可以更简单快速地进行字符串排序
	strSlice := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	sort.Sort(strSlice)
	for _, v := range strSlice {
		fmt.Println(v)
	}

	// 另外：sort 包在 sort.Interface 对各类型的封装上还有更进一步的简化，下面使用 sort.Strings 继续对代码进行简化：
	strs := []string{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	// 使用 sort.Strings 直接对字符串切片进行排序
	sort.Strings(strs)
	for _, v := range strs {
		fmt.Println(v)
	}

	// 2) 对整型切片进行排序
	// 除了字符串可以使用 sort 包进行便捷排序外，还可以使用 sort.IntSlice 进行整型切片的排序。sort.IntSlice 的定义如下：
	// type IntSlice []int
	// func (p IntSlice) Len() int           { return len(p) }
	// func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
	// func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
	// // Sort is a convenience method.
	// func (p IntSlice) Sort() { Sort(p) }
	intSlice := sort.IntSlice{4, 1, 9, 6, 3, 9}
	sort.Sort(intSlice)
	for _, v := range intSlice {
		fmt.Println(v)
	}

	// 简便方式：
	ints := []int{4, 2, 9, 6, 3, 8}
	// 直接排序方法
	sort.Ints(ints)
	for _, v := range ints {
		fmt.Println(v)
	}

	//3) 定义待排序的结构体切片
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}

	sort.Sort(heros)
	for _, v := range heros {
		fmt.Printf("%+v\n", v) // +v 打印的数据 有键名
	}
	// &{Name:关羽 Kind:1}
	// &{Name:吕布 Kind:1}
	// &{Name:李白 Kind:2}
	// &{Name:貂蝉 Kind:2}
	// &{Name:妲己 Kind:3}
	// &{Name:诸葛亮 Kind:3}

	// 4) 使用sort.Slice进行切片元素排序
	// 从 Go 1.8 开始，Go语言在 sort 包中提供了 sort.Slice() 函数进行更为简便的排序方法。
	// sort.Slice() 函数只要求传入需要排序的数据，以及一个排序时对元素的回调函数，类型为 func(i,j int)bool，定义如下：
	// func Slice(slice interface{}, less func(i, j int) bool)
	heros = Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}
	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind < heros[j].Kind
		}
		return heros[i].Name < heros[j].Name
	})
	for _, v := range heros {
		fmt.Printf("%+v\n", v) // +v 打印的数据 有键名
	}
	// &{Name:关羽 Kind:1}
	// &{Name:吕布 Kind:1}
	// &{Name:李白 Kind:2}
	// &{Name:貂蝉 Kind:2}
	// &{Name:妲己 Kind:3}
	// &{Name:诸葛亮 Kind:3}

}

// 对结构体进行排序。结构体比基本类型更为复杂，排序时不能像数值和字符串一样拥有一些固定的单一原则。
// 结构体的多个字段在排序中可能会存在多种排序的规则，
// 例如，结构体中的名字按字母升序排列，数值按从小到大的顺序排序。一般在多种规则同时存在时，需要确定规则的优先度，如先按名字排序，再按年龄排序等

// 将一批英雄名单使用结构体定义，英雄名单的结构体中定义了英雄的名字和分类。排序时要求按照英雄的分类进行排序，相同分类的情况下按名字进行排序

// HeroKind 定义为 int
type HeroKind int

// 批量定义英雄分类
const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

// Hero 英雄结构体
type Hero struct {
	Name string
	Kind HeroKind
}

// Heros 定义为 Hero指针的切片，用来存放要排序的英雄
type Heros []*Hero

// Len 实现 sort.Interface 接口的 Len() 方法
func (s Heros) Len() int {
	return len(s)
}

// Less 实现 sort.Interface 接口的 Less() 方法
func (s Heros) Less(i, j int) bool {
	// 如果英雄的分类不一致时, 优先对分类进行排序
	if s[i].Kind != s[j].Kind {
		return s[i].Kind < s[j].Kind
	}
	// 如果英雄分类一致，按英雄名称排序
	return s[i].Name < s[j].Name
}

// Swap 实现 sort.Interface 接口的 Swap 方法
func (s Heros) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
