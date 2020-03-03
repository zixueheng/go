package main

import "fmt"
import "time"

// 计时器（Timer）的原理和倒计时闹钟类似，都是给定多少时间后触发。
// 打点器（Ticker）的原理和钟表类似，钟表每到整点就会触发。
// 这两种方法创建后会返回 time.Ticker 对象和 time.Timer 对象，里面通过一个 C 成员，类型是只能接收的时间通道（<-chan Time），使用这个通道就可以获得时间触发的通知。

/**
 * ticker只要定义完成，从此刻开始计时，不需要任何其他的操作，每隔固定时间都会触发。
 * timer定时器，是到固定时间后会执行一次
 * 如果timer定时器要每隔间隔的时间执行，实现ticker的效果，使用 func (t *Timer) Reset(d Duration) bool
 */

func main() {
	// 计时器 2秒后触发
	var timer = time.NewTimer(2 * time.Second) // 返回 *time.Timer 类型变量

	// 打点器 500毫秒触发一次
	var ticker = time.NewTicker(500 * time.Millisecond) // 返回 *time.Ticker 类型变量

	var i = 0 // 用来计数

	// 不断地检查通道情况
	for {
		// 多路复用通道
		select {
		case <-ticker.C: // 打点器触发
			// 记录触发了多少次
			i++
			fmt.Printf("打点器触发%d次\n", i)
		case <-timer.C: // 计时器到时间了
			fmt.Println("计时器触发")
			goto stopS // 跳出循环
		}
	}
stopS:
	{
		fmt.Println("结束了")
	}
}

// 打点器触发1次
// 打点器触发2次
// 打点器触发3次
// 打点器触发4次
// 计时器触发
// 结束了
