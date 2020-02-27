# Goroutine 介绍
goroutine 是一种非常轻量级的实现，可在单个进程里执行成千上万的并发任务，它是Go语言并发设计的核心。
说到底 goroutine 其实就是线程，但是它比线程更小，十几个 goroutine 可能体现在底层就是五六个线程，而且Go语言内部也实现了 goroutine 之间的内存共享。

使用 go 关键字就可以创建 goroutine，将 go 声明放到一个需调用的函数之前，在相同地址空间调用运行这个函数，这样该函数执行时便会作为一个独立的并发线程，这种线程在Go语言中则被称为 goroutine。  

```go
//go 关键字放在方法调用前新建一个 goroutine 并执行方法体
go GetThingDone(param1, param2)

//新建一个匿名方法并执行
go func(param1, param2) {
}(val1, val2)

//直接新建一个 goroutine 并在 goroutine 中执行代码块
go {
    //do someting...
}
```
因为 goroutine 在多核 cpu 环境下是并行的，如果代码块在多个 goroutine 中执行，那么我们就实现了代码的并行。
如果需要了解程序的执行情况，怎么拿到并行的结果呢？需要配合使用channel进行。

# channel

channel 是Go语言在语言级别提供的 goroutine 间的通信方式。我们可以使用 channel 在两个或多个 goroutine 之间传递消息。

channel 是进程内的通信方式，因此通过 channel 传递对象的过程和调用函数时的参数传递行为比较一致，比如也可以传递指针等。如果需要跨进程通信，我们建议用分布式系统的方法来解决，比如使用 Socket 或者 HTTP 等通信协议。Go语言对于网络方面也有非常完善的支持。

channel 是类型相关的，也就是说，一个 channel 只能传递一种类型的值，这个类型需要在声明 channel 时指定。如果对 Unix 管道有所了解的话，就不难理解 channel，可以将其认为是一种类型安全的管道。

定义一个 channel 时，也需要定义发送到 channel 的值的类型，注意，必须使用 make 创建 channel，代码如下所示：
```go
ci := make(chan int)
cs := make(chan string)
cf := make(chan interface{})
```

# 示例

```go
    package main
    import (
        "fmt"
        "time"
    )
    func running() {
        var times int
        // 构建一个无限循环
        for {
            times++
            fmt.Println("tick", times)
            // 延时1秒
            time.Sleep(time.Second)
        }
    }
    // 代码执行后，命令行会不断地输出 tick，同时可以使用 fmt.Scanln() 接受用户输入。两个环节可以同时进行。
    // Go 程序在启动时，运行时（runtime）会默认为 main() 函数创建一个 goroutine。在 main() 函数的 goroutine 中执行到 go running 语句时，
    // 归属于 running() 函数的 goroutine 被创建，running() 函数开始在自己的 goroutine 中执行。此时，main() 继续执行，两个 goroutine 通过 Go 程序的调度机制同时运作。
    func main() {
        // 并发执行程序
        go running()
        // 接受命令行输入, 不做任何事情
        var input string
        fmt.Scanln(&input)
    }
```

代码执行后，命令行会不断地输出 tick，同时可以使用 fmt.Scanln() 接受用户输入。两个环节可以同时进行。

