package main

import "sync/atomic"
import "fmt"

// Go语言程序可以使用通道进行多个 goroutine 间的数据交换，但这仅仅是数据同步中的一种方法。通道内部的实现依然使用了各种锁，因此优雅代码的代价是性能。
// 在某些轻量级的场合，原子访问（atomic包）、互斥锁（sync.Mutex）以及等待组（sync.WaitGroup）能最大程度满足需求。

// 本例中只是对变量进行增减操作，虽然可以使用互斥锁（sync.Mutex）解决竞态问题，但是对性能消耗较大。在这种情况下，推荐使用原子操作（atomic）进行变量操作

var id int64

// GenID 同步的递增 seq
func GenID() int64 {
	return atomic.AddInt64(&id, 1)
}

func demo1() {
	//生成10个并发序列号
	for i := 0; i < 10; i++ {
		go GenID()
	}

	fmt.Println(GenID()) // 前面并发10生成了 seq，这里再执行应该是 9+1=10 了
}

func main() {
	demo1()
	// 注意 运行使用 go run -race main.go，不加-race结果不一样 不知道什么问题？？？
}

// Go语言包中的 sync 包提供了两种锁类型：sync.Mutex 和 sync.RWMutex。

// Mutex 是最简单的一种锁类型，同时也比较暴力，当一个 goroutine 获得了 Mutex 后，其他 goroutine 就只能乖乖等到这个 goroutine 释放该 Mutex。

// RWMutex 相对友好些，是经典的单写多读模型。在读锁占用的情况下，会阻止写，但不阻止读，也就是多个 goroutine 可同时获取读锁（调用 RLock() 方法；
// 而写锁（调用 Lock() 方法）会阻止任何其他 goroutine（无论读和写）进来，整个锁相当于由该 goroutine 独占

// 对于这两种锁类型，任何一个 Lock() 或 RLock() 均需要保证对应有 Unlock() 或 RUnlock() 调用与之对应，否则可能导致等待该锁的所有 goroutine 处于饥饿状态，甚至可能导致死锁。

var (
	// 逻辑中使用的某个变量
	count int
	// 与变量对应的使用互斥锁
	countGuard sync.Mutex // 一般情况下，建议将互斥锁的粒度设置得越小越好，降低因为共享访问时等待的时间。采用格式：变量名+Guard，表示这个互斥锁用于保护这个变量
)

// GetCount 安全获取遍历
func GetCount() int {
	// 锁定
	countGuard.Lock()
	// 在函数退出时解除锁定
	defer countGuard.Unlock()
	return count
}

// SetCount 安全的设置变量
func SetCount(c int) {
	countGuard.Lock()
	count = c
	countGuard.Unlock()
}

func demo2() {
	// 可以进行并发安全的设置
	SetCount(1)
	// 可以进行并发安全的获取
	fmt.Println(GetCount())
}

// 3、在读多写少的环境中，可以优先使用读写互斥锁（sync.RWMutex），它比互斥锁更加高效。sync 包中的 RWMutex 提供了读写互斥锁的封装。

var (
	// 逻辑中使用的某个变量
	count2 int
	// 与变量对应的使用互斥锁
	count2Guard sync.RWMutex
)

// GetCount2 获取变量，在没有写的情况下可以并发的读，性能较好
func GetCount2() int {
	// 锁定
	count2Guard.RLock()
	// 在函数退出时解除锁定
	defer count2Guard.RUnlock()
	return count
}

// SetCount2 写变量，会阻塞其他所有 goroutine
func SetCount2(c int) {
	count2Guard.Lock()
	count = c
	count2Guard.Unclock()
}

func demo3() {
	SetCount2(2)
	fmt.Println(GetCount2())
}
