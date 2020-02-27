package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var (
	counter int64
	wg      sync.WaitGroup
)

// WaitGroup 对象内部有一个计数器，最初从0开始，它有三个方法：Add(), Done(), Wait() 用来控制计数器的数量。
// Add(n) 把计数器设置为n ，Done() 每次把计数器-1 ，wait() 会阻塞代码的运行，直到计数器地值减为0。

// 原子函数 atomic 能够以很底层的加锁机制来同步访问整型变量和指针
// 使用了 atmoic 包的 AddInt64 函数，这个函数会同步整型值的加法，方法是强制同一时刻只能有一个 gorountie 运行并完成这个加法操作。
// 当 goroutine 试图去调用任何原子函数时，这些 goroutine 都会自动根据所引用的变量做同步处理
func incCounter(id int) {
	defer wg.Done() // 函数执行一次后计数器减1
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter, 1) // 安全的对counter加1
		runtime.Gosched()
	}
}

func demo1() {
	wg.Add(2) // 设置计数器从2开始
	go incCounter(1)
	go incCounter(2)
	wg.Wait()            // 等待goroutine结束, 直到计数器地值减为0
	fmt.Println(counter) // 等计算器值为0时 这行代码才会执行（因为上一行 Wait 已阻塞执行）
}

func main() {
	// demo1() // 4
	// demo2()
	demo3()
}

var shutdown int64

// 另外两个有用的原子函数是 LoadInt64 和 StoreInt64。
// 这两个函数提供了一种安全地读和写一个整型值的方式。
// 下面是代码就使用了 LoadInt64 和 StoreInt64 函数来创建一个同步标志，这个标志可以向程序里多个 goroutine 通知某个特殊状态
func demo2() {
	wg.Add(2)
	go doWork("A")
	go doWork("B")
	time.Sleep(1 * time.Second) // 等待1秒 让上面的 goroutine 执行几遍
	fmt.Println("Shutdown Now")
	// 使用 StoreInt64 函数来安全地修改 shutdown 变量的值,
	// 如果哪个 doWork goroutine 试图在 main 函数调用 StoreInt64 的同时调用 LoadInt64 函数，那么原子函数会将这些调用互相同步，保证这些操作都是安全的，不会进入竞争状态。
	atomic.StoreInt64(&shutdown, 1) // 1秒后 将 shutdown 置为1
	wg.Wait()                       // 阻塞执行，不让demo2函数结束直到 wg 计数器为0
}

func doWork(name string) {
	defer wg.Done()
	for {
		fmt.Printf("Doing %s Work\n", name)
		time.Sleep(250 * time.Millisecond) // 模拟此 任务需要执行 250毫秒
		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s Down, shutdown=%d\n\n", name, atomic.LoadInt64(&shutdown))
			break
		} else {
			fmt.Printf("shutdown=%d\n", atomic.LoadInt64(&shutdown))
		}
	}
}

var mutex sync.Mutex

// 互斥锁
// 另一种同步访问共享资源的方式是使用互斥锁，互斥锁这个名字来自互斥的概念。互斥锁用于在代码上创建一个临界区，保证同一时间只有一个 goroutine 可以执行这个临界代码。
func demo3() {
	wg.Add(2)
	go incCounter2(1)
	go incCounter2(2)
	wg.Wait()
	fmt.Println(counter)
}

// 同一时刻只有一个 goroutine 可以进入临界区。
// 之后直到调用 Unlock 函数之后，其他 goroutine 才能进去临界区。
// 当调用 runtime.Gosched 函数强制将当前 goroutine 退出当前线程后，调度器会再次分配这个 goroutine 继续运行。
func incCounter2(id int) {
	defer wg.Done()
	for count := 0; count < 2; count++ {
		//同一时刻只允许一个goroutine进入这个临界区
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock() //释放锁，允许其他正在等待的goroutine进入临界区
	}
}
