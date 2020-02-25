# Go 命令行参数解析 flag包

在编写命令行程序（工具、server）时，我们有时需要对命令参数进行解析，各种编程语言一般都会提供解析命令行参数的方法或库，以方便程序员使用。在Go语言中的 flag 包中，提供了命令行参数解析的功能。

这里介绍几个概念：
    命令行参数（或参数）：是指运行程序时提供的参数；
    已定义命令行参数：是指程序中通过 flag.Type 这种形式定义了的参数；
    非 flag（non-flag）命令行参数（或保留的命令行参数）：可以简单理解为 flag 包不能解析的参数。  
## 有以下两种常用的定义命令行 flag 参数的方法：

### 1、flag.Type(flag 名, 默认值, 帮助信息) *Type
Type 可以是 Int、String、Bool 等，返回值为一个相应类型的指针  
```go
name := flag.String("name", "某某某", "姓名") // 返回一个 string 类型的指针 *string
age := flag.Int("age", 18, "年龄")
married := flag.Bool("married", false, "婚否")
delay := flag.Duration("d", 0, "时间间隔")
```
### 2、flag.TypeVar(Type 指针, flag 名, 默认值, 帮助信息)
TypeVar 可以是 IntVar、StringVar、BoolVar 等，其功能为将 flag 绑定到一个变量上  
```go
var name string
var age int
var married bool
var delay time.Duration
flag.StringVar(&name, "name", "张三", "姓名")
flag.IntVar(&age, "age", 18, "年龄")
flag.BoolVar(&married, "married", false, "婚否")
flag.DurationVar(&delay, "d", 0, "时间间隔")
```

## flag.Parse()
通过以上两种方法定义好命令行 flag 参数后，需要通过调用 flag.Parse() 来对命令行参数进行解析。
支持的命令行参数格式有以下几种：
    -flag：只支持 bool 类型；
    -flag=x；
    -flag x：只支持非 bool 类型。
其中，布尔类型的参数必须使用等号的方式指定。

flag 包的其他函数：
```go
flag.Args()  //返回命令行参数后的其他参数，以 []string 类型
flag.NArg()  //返回命令行参数后的其他参数个数
flag.NFlag() //返回使用的命令行参 数个数
```
