# Channel 通道

## 通道的特性
Go语言中的通道（channel）是一种特殊的类型。在任何时候，同时只能有一个 goroutine 访问通道进行发送和获取数据。goroutine 间通过通道就可以通信。

通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。

## 声明通道类型
通道本身需要一个类型进行修饰，就像切片类型需要标识元素类型。通道的元素类型就是在其内部传输的数据类型，声明如下：

var 通道变量 chan 通道类型
- 通道类型：通道内的数据类型。
- 通道变量：保存通道的变量。

chan 类型的空值是 nil，声明后需要配合 make 后才能使用。

## 创建通道
通道是引用类型，需要使用 make 进行创建，格式如下：

通道实例 := make(chan 数据类型)
- 数据类型：通道内传输的元素类型。
- 通道实例：通过make创建的通道句柄。

请看下面的例子： 
```go
ch1 := make(chan int)                 // 创建一个整型类型的通道
ch2 := make(chan interface{})         // 创建一个空接口类型的通道, 可以存放任意格式
type Equip struct{ /* 一些字段 */ }
ch2 := make(chan *Equip)             // 创建Equip指针类型的通道, 可以存放*Equip
```

## 使用通道发送数据
通道创建后，就可以使用通道进行发送和接收操作。

#### 1) 通道发送数据的格式
通道的发送使用特殊的操作符`<-`，将数据通过通道发送的格式为：
`通道变量 <- 值`
- 通道变量：通过make创建好的通道实例。
- 值：可以是变量、常量、表达式或者函数返回值等。值的类型必须与ch通道的元素类型一致。

#### 2) 通过通道发送数据的例子
使用 make 创建一个通道后，就可以使用`<-`向通道发送数据，代码如下：
```go
// 创建一个空接口通道
ch := make(chan interface{})
// 将0放入通道中
ch <- 0
// 将hello字符串放入通道中
ch <- "hello"
```
###  3) 发送将持续阻塞直到数据被接收
把数据往通道中发送时，如果接收方一直都没有接收，那么发送操作将持续阻塞。Go 程序运行时能智能地发现一些永远无法发送成功的语句并做出提示，代码如下： 
```go
package main
func main() {
    // 创建一个整型通道
    ch := make(chan int)
    // 尝试将0通过通道发送
    ch <- 0
}
```
运行代码，报错：
`fatal error: all goroutines are asleep - deadlock!`
报错的意思是：运行时发现所有的 goroutine（包括main）都处于等待 goroutine。也就是说所有 goroutine 中的 channel 并没有形成发送和接收对应的代码。 

## 使用通道接收数据

通道接收同样使用`<-`操作符，通道接收有如下特性：
① 通道的收发操作在不同的两个 goroutine 间进行。
由于通道的数据在没有接收方处理时，数据发送方会持续阻塞，因此通道的接收必定在另外一个 goroutine 中进行。

② 接收将持续阻塞直到发送方发送数据。
如果接收方接收时，通道中没有发送方发送数据，接收方也会发生阻塞，直到发送方发送数据为止。

③ 每次接收一个元素。
通道一次只能接收一个数据元素。

通道的数据接收一共有以下 4 种写法。

#### 1) 阻塞接收数据
阻塞模式接收数据时，将接收变量作为`<-`操作符的左值，格式如下：
`data := <-ch`
执行该语句时将会阻塞，直到接收到数据并赋值给 data 变量。

#### 2) 非阻塞接收数据
使用非阻塞方式从通道接收数据时，语句不会发生阻塞，格式如下：
`data, ok := <-ch`

- data：表示接收到的数据。未接收到数据时，data 为通道类型的零值。
- ok：表示是否接收到数据。

非阻塞的通道接收方法可能造成高的 CPU 占用，因此使用非常少。如果需要实现接收超时检测，可以配合 select 和计时器 channel 进行，可以参见后面的内容。

#### 3) 接收任意数据，忽略接收的数据
阻塞接收数据后，忽略从通道返回的数据，格式如下：
`<-ch`
执行该语句时将会发生阻塞，直到接收到数据，但接收到的数据会被忽略。这个方式实际上只是通过通道在 goroutine 间阻塞收发实现并发同步。

使用通道做并发同步的写法，可以参考下面的例子：

```go
package main
import (
    "fmt"
)
func main() {
    // 构建一个通道
    ch := make(chan int)
    // 开启一个并发匿名函数
    go func() {
        fmt.Println("start goroutine")
        // 通过通道通知main的goroutine
        ch <- 0
        fmt.Println("exit goroutine")
    }()
    fmt.Println("wait goroutine")
    // 等待匿名goroutine
    <-ch
    fmt.Println("all done")
}
```

#### 4) 循环接收
通道的数据接收可以借用 `for range` 语句进行多个元素的接收操作，格式如下：
```go
for data := range ch {
}
```

通道 ch 是可以进行遍历的，遍历的结果就是接收到的数据。数据类型就是通道的数据类型。通过 for 遍历获得的变量只有一个，即上面例子中的 data。

使用 for 从通道中接收数据：
```go
package main
import (
    "fmt"
    "time"
)
func main() {
    // 构建一个通道
    ch := make(chan int)
    // 开启一个并发匿名函数
    go func() {
        // 从3循环到0
        for i := 3; i >= 0; i-- {
            // 发送3到0之间的数值
            ch <- i
            // 每次发送完时等待
            time.Sleep(time.Second)
        }
    }()
    // 遍历接收通道数据
    for data := range ch {
        // 打印通道数据
        fmt.Println(data)
        // 当遇到数据0时, 退出接收循环
        if data == 0 {
                break
        }
    }
}
```

## 单向通道的声明格式
我们在将一个 channel 变量传递到一个函数时，可以通过将其指定为单向 channel 变量，从而限制该函数中可以对此 channel 的操作，比如只能往这个 channel 写，或者只能从这个 channel 读。

单向 channel 变量的声明非常简单，只能发送的通道类型为`chan<-`，只能接收的通道类型为`<-chan`，格式如下：
`var 通道实例 chan<- 元素类型   // 只能发送通道`
`var 通道实例 <-chan 元素类型   // 只能接收通道`

- 元素类型：通道包含的元素类型。
- 通道实例：声明的通道变量。

```go
ch := make(chan int)
// 声明一个只能发送的通道类型, 并赋值为ch
var chSendOnly chan<- int = ch
//声明一个只能接收的通道类型, 并赋值为ch
var chRecvOnly <-chan int = ch
```
上面的例子中，chSendOnly 只能发送数据，如果尝试接收数据，将会出现如下报错：
`invalid operation: <-chSendOnly (receive from send-only type chan<- int)`
同理，chRecvOnly 也是不能发送的。

## 关闭 channel
关闭 channel 非常简单，直接使用 Go语言内置的 close() 函数即可：
`close(ch)`

在介绍了如何关闭 channel 之后，我们就多了一个问题：如何判断一个 channel 是否已经被关闭？我们可以在读取的时候使用多重返回值的方式：
`x, ok := <-ch`

这个用法与 map 中的按键获取 value 的过程比较类似，只需要看第二个 bool 返回值即可，如果返回值是 false 则表示 ch 已经被关闭。

# Go语言无缓冲的通道

Go语言中无缓冲的通道（unbuffered channel）是指在接收前没有能力保存任何值的通道。这种类型的通道要求发送 goroutine 和接收 goroutine 同时准备好，才能完成发送和接收操作。

 如果两个 goroutine 没有同时准备好，通道会导致先执行发送或接收操作的 goroutine 阻塞等待。这种对通道进行发送和接收的交互行为本身就是同步的。其中任意一个操作都无法离开另一个操作单独存在。

 阻塞指的是由于某种原因数据没有到达，当前协程（线程）持续处于等待状态，直到条件满足才解除阻塞。

 同步指的是在两个或多个协程（线程）之间，保持数据内容一致性的机制。

【示例 1】在网球比赛中，两位选手会把球在两个人之间来回传递。选手总是处在以下两种状态之一，要么在等待接球，要么将球打向对方。可以使用两个 goroutine 来模拟网球比赛，并使用无缓冲的通道来模拟球的来回，代码如下所示。

```go
    // 这个示例程序展示如何用无缓冲的通道来模拟
    // 2 个goroutine 间的网球比赛
    package main
    import (
        "fmt"
        "math/rand"
        "sync"
        "time"
    )
    // wg 用来等待程序结束
    var wg sync.WaitGroup
    func init() {
        rand.Seed(time.Now().UnixNano())
    }
    // main 是所有Go 程序的入口
    func main() {
        // 创建一个无缓冲的通道
        court := make(chan int)
        // 计数加 2，表示要等待两个goroutine
        wg.Add(2)
        // 启动两个选手
        go player("Nadal", court)
        go player("Djokovic", court)
        // 发球
        court <- 1
        // 等待游戏结束
        wg.Wait()
    }
    // player 模拟一个选手在打网球
    func player(name string, court chan int) {
        // 在函数退出时调用Done 来通知main 函数工作已经完成
        defer wg.Done()
        for {
            // 等待球被击打过来
            ball, ok := <-court
            if !ok {
                // 如果通道被关闭，我们就赢了
                fmt.Printf("Player %s Won\n", name)
                return
            }
            // 选随机数，然后用这个数来判断我们是否丢球
            n := rand.Intn(100)
            if n%13 == 0 {
                fmt.Printf("Player %s Missed\n", name)
                // 关闭通道，表示我们输了
                close(court)
                return
            }
            // 显示击球数，并将击球数加1
            fmt.Printf("Player %s Hit %d\n", name, ball)
            ball++
            // 将球打向对手
            court <- ball
        }
    }
```

