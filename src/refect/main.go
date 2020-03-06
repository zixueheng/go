package main

import (
	"fmt"
	"reflect"
)

// 反射

// Enum 定义为 int
type Enum int

// Cat 结构体
type Cat struct{}

func demo1() {
	var zero Enum = 0
	var typeOfZero reflect.Type = reflect.TypeOf(zero) // reflect.Type 是一个接口包含 Name() Kind()等方法
	fmt.Println(typeOfZero.Name(), typeOfZero.Kind())
	// Enum int
	// 可以理解为 zero 的类型名称是Enum，实际类型是 int

	var cat Cat = Cat{}
	var typeOfCat reflect.Type = reflect.TypeOf(cat)
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
	// Cat struct
	// 可以理解为 cat 类型名称是Cat，实际类型是 struct

	// 这里试下指针
	var catPtr *Cat = &Cat{}
	var typeOfCatPtr reflect.Type = reflect.TypeOf(catPtr)
	fmt.Println(typeOfCatPtr.Name(), typeOfCatPtr.Kind())
	// "" ptr
	// 指针没有名字返回是空字符串，大类Kind ptr

	// 对指针获取反射对象时，可以通过 reflect.Elem() 方法获取这个指针指向的元素类型
	// 这个获取过程被称为取元素，等效于对指针类型变量做了一个*操作
	var typeOfCatPtr2 reflect.Type = typeOfCatPtr.Elem()    // Elem() 返回该类型的元素类型(这里理解返回实际类型)，如果该类型的Kind不是Array、Chan、Map、Ptr或Slice，会panic
	fmt.Println(typeOfCatPtr2.Name(), typeOfCatPtr2.Kind()) // 输出指针变量指向元素的类型名称和种类，得到了 cat 的类型名称（cat）和种类（struct）
	// Cat struct
}

// 使用反射获取结构体的成员类型
func demo2() {
	// 声明一个空结构体
	type cat struct {
		Name string
		// 带有结构体tag的字段
		Type int `json:"type" id:"100"`
	}
	// 创建cat的实例
	ins := cat{Name: "mimi", Type: 1}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	// 遍历结构体所有成员
	for i := 0; i < typeOfCat.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := typeOfCat.Field(i)
		// 输出成员名和tag
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}
	// 通过字段名, 找到字段类型信息
	if catType, ok := typeOfCat.FieldByName("Type"); ok {
		// 从tag中取出需要的tag
		fmt.Println(catType.Tag.Get("json"), catType.Tag.Get("id"))
	}
}

// reflect.Value
func demo3() {
	var w uint8 = 'x'
	var valueOfW reflect.Value = reflect.ValueOf(w) // reflect.Value 是一个结构体
	fmt.Println("w Type: ", valueOfW.Type())        // 返回valueOfW持有的值的类型的Type表示
	// w Type:  uint8

	fmt.Println("kind is uint8: ", valueOfW.Kind() == reflect.Uint8)
	// kind is uint8:  true

	// func (v Value) Uint() uint64
	// 返回v持有的无符号整数（表示为uint64），如v的Kind不是Uint、Uintptr、Uint8、Uint16、Uint32、Uint64会panic
	x := uint8(valueOfW.Uint())
	fmt.Println(x)
	// 120

}

// 结构体字段值修改
func demo4() {
	// T 结构体字段名要大写，是因为结构体中只有可导出的字段是“可设置”的。
	type T struct {
		A int
		B string
	}
	t := T{23, "skidoo"}

	// &t 表示要传t地址指针 进去而不是传拷贝值（如果不传地址进去反射对象存储的就不是原始值）
	p := reflect.ValueOf(&t)
	// 再调用 Elem() 因为我们实际想修改 指针 所指向的 值，调用 Elem()方法能够对指针进行“解引用”，然后将结果存储到反射 Value 类型对象 s 中
	s := p.Elem()
	fmt.Println(s.CanSet()) // true

	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	fmt.Println("t is now", t)
	// t is now {77 Sunset Strip}
}

// 值的修改从表面意义上叫可寻址，换一种说法就是值必须“可被设置”。那么，想修改变量值，一般的步骤是：
//     取这个变量的地址或者这个变量所在的结构体已经是指针类型。
//     使用 reflect.ValueOf 进行值包装。
//     通过 Value.Elem() 获得指针值指向的元素值对象（Value），因为值对象（Value）内部对象为指针时，使用 set 设置时会报出宕机错误。
//     使用 Value.Set 设置值。

func main() {
	// demo1()

	// demo2()
	// name: Name  tag: ''
	// name: Type  tag: 'json:"type" id:"100"'
	// type 100

	// demo3()

	// demo4()

	// demo5()

	demo6()
}

// 对指针获取反射对象时，可以通过 reflect.Elem() 方法获取这个指针指向的元素类型。这个获取过程被称为取元素，等效于对指针类型变量做了一个*操作
func demo5() {
	// 声明一个空结构体
	type cat struct{}
	// 创建cat的实例
	ins := &cat{}
	// 获取结构体实例的反射类型对象
	typeOfCat := reflect.TypeOf(ins)
	// 显示反射类型对象的名称和种类
	fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind())
	// name: ''  kind: 'ptr'
	// 取类型的元素
	typeOfCat = typeOfCat.Elem()
	// 显示反射类型对象的名称和种类
	fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())
	// element name: 'cat', element kind: 'struct'
}

// 当已知 reflect.Type 时，可以动态地创建这个类型的实例，实例的类型为指针。例如 reflect.Type 的类型为 int 时，创建 int 的指针，即*int
func demo6() {
	var a int = 10
	// 取变量a的反射类型对象
	var typeOfA reflect.Type = reflect.TypeOf(a)

	// 根据反射类型对象创建类型实例
	var aIns reflect.Value = reflect.New(typeOfA)
	// 输出Value的类型和种类
	fmt.Println(aIns.Type(), aIns.Kind())
	// *int ptr

	fmt.Println(aIns.Elem()) // 查看实际的值，新生成的都是零值
	// 0

	aIns.Elem().SetInt(11)   // 赋值
	fmt.Println(aIns.Elem()) // 11

	// *aIns = 11
	// fmt.Println(*aIns)
}
