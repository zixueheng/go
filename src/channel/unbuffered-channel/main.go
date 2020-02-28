package main

// Go语言中无缓冲的通道（unbuffered channel）是指在接收前没有能力保存任何值的通道。这种类型的通道要求发送 goroutine 和接收 goroutine 同时准备好，才能完成发送和接收操作。

// 如果两个 goroutine 没有同时准备好，通道会导致先执行发送或接收操作的 goroutine 阻塞等待。这种对通道进行发送和接收的交互行为本身就是同步的。其中任意一个操作都无法离开另一个操作单独存在。

// 阻塞指的是由于某种原因数据没有到达，当前协程（线程）持续处于等待状态，直到条件满足才解除阻塞。

// 同步指的是在两个或多个协程（线程）之间，保持数据内容一致性的机制。

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// wg 用来等待程序结束
var wg sync.WaitGroup

func init() {
	// 如果不使用rand.Seed(seed int64)，每次运行，得到的随机数会一样
	// 每次运行时rand.Seed(seed int64)，seed的值要不一样，这样生成的随机数才会和上次运行时生成的随机数不一样
	// 如果不执行 rand.Intn(n int) 会得到一样的值
	rand.Seed(time.Now().UnixNano()) // 随机种子
}

// main 是所有Go 程序的入口
func main() {
	// demo1()
	// 选手 2号 Hit 1
	// 选手 1号 Hit 2
	// 选手 2号 Hit 3
	// 选手 1号 Hit 4
	// 选手 2号 Hit 5
	// 选手 1号 Hit 6
	// 选手 2号 Hit 7
	// 选手 1号 Hit 8
	// 选手 2号 Missed
	// 选手 1号 Won

	demo2()
	// 运动员1持有接力棒
	// 运动员2准备
	// 运动员1将接力棒交于远动员2
	// 运动员2持有接力棒
	// 运动员3准备
	// 运动员2将接力棒交于远动员3
	// 运动员3持有接力棒
	// 运动员4准备
	// 运动员3将接力棒交于远动员4
	// 运动员4持有接力棒
	// 运动员4完成跑步, 比赛结束
}

// demo1 展示如何用无缓冲的通道来模拟 2个goroutine 间的网球比赛
func demo1() {
	// 创建一个无缓冲的通道，让两个 goroutine 在击球时能够互相同步
	court := make(chan int)

	// 计数加 2，表示要等待两个goroutine
	wg.Add(2)

	// 创建了参与比赛的两个 goroutine，在这个时候，两个 goroutine 都阻塞住等待击球
	go player("1号", court)
	go player("2号", court)

	// 发球
	court <- 1 // 将球发到通道里，程序开始执行这个比赛，直到某个 goroutine 输掉比赛

	// 等待游戏结束
	wg.Wait()
}

// player 模拟一个选手在打网球
func player(name string, court chan int) {
	// 在函数退出时调用Done 来通知main 函数工作已经完成
	defer wg.Done()

	// 无限循环的 for 语句，在这个循环里，是玩游戏的过程
	for {
		// 等待球被击打过来
		ball, ok := <-court // goroutine 从通道接收数据，用来表示等待接球。这个接收动作会锁住 goroutine，直到有数据发送到通道里。通道的接收动作返回时

		if !ok {
			// 如果通道被关闭（即没有球发过来），我们就赢了
			fmt.Printf("选手 %s Won\n", name)
			return
		}

		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100) // 注意 要用 rand.Seed(time.Now().UnixNano()) 添加随机种子（在 init函数 里面写），保证 rand出来的值是随机的
		if n%13 == 0 {
			fmt.Printf("选手 %s Missed\n", name)
			// 关闭通道，表示我们输了
			close(court)
			return
		}

		// 显示击球数，并将击球数加1
		fmt.Printf("选手 %s Hit %d\n", name, ball)
		ball++

		// 将球打向对手
		court <- ball // 将 ball 作为球重新放入通道，发送给另一位选手。在这个时刻，两个 goroutine 都会被锁住，直到交换完成
	}
}

// demo2 用不同的模式，使用无缓冲的通道，在 goroutine 之间同步数据，来模拟接力比赛。在接力比赛里，4 个跑步者围绕赛道轮流跑。
// 第二个、第三个和第四个跑步者要接到前一位跑步者的接力棒后才能起跑。比赛中最重要的部分是要传递接力棒，要求同步传递。
// 在同步接力棒的时候，参与接力的两个跑步者必须在同一时刻准备好交接
func demo2() {
	// 创建了一个无缓冲的 int 类型的通道 baton，用来同步传递接力棒
	baton := make(chan int)

	// 给 WaitGroup 加 1，这样 main 函数就会等最后一位跑步者跑步结束
	wg.Add(1)

	// 第一位跑步者持有接力棒
	go Runner(baton)

	// 将接力棒交给这个跑步者，比赛开始
	baton <- 1

	// demo2函数 阻塞在 WaitGroup，等候最后一位跑步者完成比赛
	wg.Wait()
}

// Runner 模拟接力比赛中的一位跑步者
func Runner(baton chan int) {
	var newRunner int

	// goroutine 对 baton 通道执行接收操作，表示等候接力棒
	runner := <-baton

	// 开始绕着跑道跑步
	fmt.Printf("运动员%d持有接力棒\n", runner)

	// 一旦接力棒传了进来，就会创建一位新跑步者，准备接力下一棒，直到 goroutine 是第4个跑步者
	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("运动员%d准备\n", newRunner)
		go Runner(baton)
	}

	// 跑步者围绕跑道跑 1 s
	time.Sleep(1 * time.Second)

	// 如果第四个跑步者完成了比赛，就调用 Done，将 WaitGroup 减 1，之后 goroutine 返回
	if runner == 4 {
		fmt.Printf("运动员%d完成跑步, 比赛结束\n", runner)
		wg.Done()
		return
	}

	// 将接力棒交给下一位跑步者
	fmt.Printf("运动员%d将接力棒交于远动员%d\n", runner, newRunner)
	baton <- newRunner // 如果这个 goroutine 不是第四个跑步者，接力棒会交到下一个已经在等待的跑步者手上。在这个时候，goroutine 会被锁住，直到交接完成
}
