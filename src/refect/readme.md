# 反射
Go语言中的反射是由 reflect 包提供支持的，它定义了两个重要的类型 Type 和 Value 任意接口值在反射中都可以理解为由 reflect.Type 和 reflect.Value 两部分组成，并且 reflect 包提供了 reflect.TypeOf() 和 reflect.ValueOf() 两个函数来获取任意对象的 Value 和 Type。



## 反射的类型对象（reflect.Type）

在Go语言程序中，使用 reflect.TypeOf() 函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息，下面通过示例来理解获取类型对象的过程：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        var a int
        typeOfA := reflect.TypeOf(a)
        fmt.Println(typeOfA.Name(), typeOfA.Kind()) // 取到typeOfA变量的类型名为int，种类（Kind）为int
        // int int
    }
```

## 反射的类型（Type）与种类（Kind）

在使用反射时，需要首先理解类型（Type）和种类（Kind）的区别。编程中，使用最多的是类型，但在反射中，当需要区分一个大品种的类型时，就会用到种类（Kind）。例如需要统一判断类型中的指针时，使用种类（Kind）信息就较为方便。

#### 1) 反射种类（Kind）的定义

Go语言程序中的类型（Type）指的是系统原生数据类型，如 int、string、bool、float32 等类型，以及使用 type  关键字定义的类型，这些类型的名称就是其类型本身的名称。例如使用 type A struct{} 定义结构体时，A 就是 struct{}  的类型。

 种类（Kind）指的是对象归属的品种，在 reflect 包中有如下定义：

```go
    type Kind uint
    const (
        Invalid Kind = iota  // 非法类型
        Bool                 // 布尔型
        Int                  // 有符号整型
        Int8                 // 有符号8位整型
        Int16                // 有符号16位整型
        Int32                // 有符号32位整型
        Int64                // 有符号64位整型
        Uint                 // 无符号整型
        Uint8                // 无符号8位整型
        Uint16               // 无符号16位整型
        Uint32               // 无符号32位整型
        Uint64               // 无符号64位整型
        Uintptr              // 指针
        Float32              // 单精度浮点数
        Float64              // 双精度浮点数
        Complex64            // 64位复数类型
        Complex128           // 128位复数类型
        Array                // 数组
        Chan                 // 通道
        Func                 // 函数
        Interface            // 接口
        Map                  // 映射
        Ptr                  // 指针
        Slice                // 切片
        String               // 字符串
        Struct               // 结构体
        UnsafePointer        // 底层指针
    )
```

Map、Slice、Chan 属于引用类型，使用起来类似于指针，但是在种类常量定义中仍然属于独立的种类，不属于 Ptr。type A struct{} 定义的结构体属于 Struct 种类，*A 属于 Ptr。

#### 2) 从类型对象中获取类型名称和种类

Go语言中的类型名称对应的反射获取方法是 reflect.Type 中的 Name() 方法，返回表示类型名称的字符串；类型归属的种类（Kind）使用的是 reflect.Type 中的 Kind() 方法，返回 reflect.Kind 类型的常量。

```go
package main

import (
	"fmt"
	"reflect"
)

// Enum 定义为 int
type Enum int

// Cat 结构体
type Cat struct{}

func main() {
	var zero Enum = 0
	var typeOfZero reflect.Type = reflect.TypeOf(zero) // reflect.Type 是一个接口包含 Name() Kind()等方法
	fmt.Println(typeOfZero.Name(), typeOfZero.Kind())
	// Enum int
	// 可以理解为 zero 的类型名称是Enum，实际类型是 int

	var cat Cat = Cat{}
	typeOfCat := reflect.TypeOf(cat)
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())
	// Cat struct
	// 可以理解为 cat 类型名称是Cat，实际类型是 struct

	var catPtr *Cat = &Cat{} // 这里试下指针
	typeOfCatPtr := reflect.TypeOf(catPtr)
	fmt.Println(typeOfCatPtr.Name(), typeOfCatPtr.Kind())
	// "" ptr
	// 指针没有名字，大类Kind ptr
}

```

## 指针与指针指向的元素

Go语言程序中对指针获取反射对象时，可以通过 reflect.Elem() 方法获取这个指针指向的元素类型，这个获取过程被称为取元素，等效于对指针类型变量做了一个`*`操作，代码如下：

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
        // 声明一个空结构体
        type cat struct {
        }
        // 创建cat的实例
        ins := &cat{}
        // 获取结构体实例的反射类型对象
        typeOfCat := reflect.TypeOf(ins)
        // 显示反射类型对象的名称和种类
        fmt.Printf("name:'%v' kind:'%v'\n", typeOfCat.Name(), typeOfCat.Kind())
        // 取类型的元素
        typeOfCat = typeOfCat.Elem()
        // 显示反射类型对象的名称和种类
        fmt.Printf("element name: '%v', element kind: '%v'\n", typeOfCat.Name(), typeOfCat.Kind())
    }
// name:'' kind:'ptr'
// element name: 'cat', element kind: 'struct'
```

## 使用反射获取结构体的成员类型

任意值通过 reflect.TypeOf() 获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象 reflect.Type 的 NumField() 和 Field() 方法获得结构体成员的详细信息。

 与成员获取相关的 reflect.Type 的方法如下表所示。

| 方法                                                        | 说明                                                         |
| ----------------------------------------------------------- | ------------------------------------------------------------ |
| Field(i int) StructField                                    | 根据索引返回索引对应的结构体字段的信息，当值不是结构体或索引超界时发生宕机 |
| NumField() int                                              | 返回结构体成员字段数量，当类型不是结构体或索引超界时发生宕机 |
| FieldByName(name string) (StructField, bool)                | 根据给定字符串返回字符串对应的结构体字段的信息，没有找到时 bool 返回 false，当类型不是结构体或索引超界时发生宕机 |
| FieldByIndex(index []int) StructField                       | 多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息，没有找到时返回零值。当类型不是结构体或索引超界时发生宕机 |
| FieldByNameFunc(match func(string) bool) (StructField,bool) | 根据匹配函数匹配需要的字段，当值不是结构体或索引超界时发生宕机 |

#### 1) 结构体字段类型

reflect.Type 的 Field() 方法返回 StructField  结构，这个结构描述结构体的成员信息，通过这个信息可以获取成员与结构体的关系，如偏移、索引、是否为匿名字段、结构体标签（StructTag）等，而且还可以通过 StructField 的 Type 字段进一步获取结构体成员的类型信息。

 StructField 的结构如下：

```go
    type StructField struct {
        Name string          // 字段名
        PkgPath string       // 字段路径
        Type      Type       // 字段反射类型对象
        Tag       StructTag  // 字段的结构体标签
        Offset    uintptr    // 字段在结构体中的相对偏移
        Index     []int      // Type.FieldByIndex中的返回的索引值
        Anonymous bool       // 是否为匿名字段
    }
```

#### 2) 获取成员反射信息

下面代码中，实例化一个结构体并遍历其结构体成员，再通过 reflect.Type 的 FieldByName() 方法查找结构体中指定名称的字段，直接获取其类型信息。

```go
    package main
    import (
        "fmt"
        "reflect"
    )
    func main() {
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
```

## 结构体标签（Struct Tag）

通过 reflect.Type 获取结构体成员信息 reflect.StructField 结构中的 Tag 被称为结构体标签（StructTag）。结构体标签是对结构体字段的额外信息标签。

 JSON、BSON 等格式进行序列化及对象关系映射（Object Relational Mapping，简称  ORM）系统都会用到结构体标签，这些系统使用标签设定字段在处理时应该具备的特殊属性和可能发生的行为。这些信息都是静态的，无须实例化结构体，可以通过反射获取到。

#### 1) 结构体标签的格式

Tag 在结构体字段后方书写的格式如下：

`key1:"value1" key2:"value2"`

结构体标签由一个或多个键值对组成；键与值使用冒号分隔，值用双引号括起来；键值对之间使用一个空格分隔。

#### 2) 从结构体标签中获取值

StructTag 拥有一些方法，可以进行 Tag 信息的解析和提取，如下所示：

- `func (tag StructTag) Get(key string) string`：根据 Tag 中的键获取对应的值，例如`key1:"value1" key2:"value2"`的 Tag 中，可以传入“key1”获得“value1”。

- `func (tag StructTag) Lookup(key string) (value string, ok bool)`：根据 Tag 中的键，查询值是否存在。

#### 3) 结构体标签格式错误导致的问题

编写 Tag 时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误。