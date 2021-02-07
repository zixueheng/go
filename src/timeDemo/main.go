package main

import (
	"fmt"
	"time"
)

// TimeFormat 格式化时间戳
func TimeFormat(timstamp int64) string {
	timeObj := time.Unix(timstamp, 0)
	// return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d", timeObj.Year(), timeObj.Month(), timeObj.Day(), timeObj.Hour(), timeObj.Minute(), timeObj.Second()) // %02d 表示 2是宽度，如果整数不够2列就补上0
	// 简单的方式
	return timeObj.Format("2006-01-02 15:04:05")
}

func main() {
	now := time.Now()
	fmt.Println(now)         // 2020-02-05 17:00:17.8896148 +0800 CST m=+0.009970701
	fmt.Println(now.Date())  // 2020 February 5
	fmt.Println(now.Year())  // 2020
	fmt.Println(now.Month()) // February
	fmt.Println(now.Day())   // 5

	fmt.Println(now.Unix()) // 1580893480 时间戳

	fmt.Println(TimeFormat(1580893480)) // 2020-02-05 17:04:40

	// 24小时后
	fmt.Println(now.Add(24 * time.Hour))

	// 定时器
	// timer := time.Tick(2 * time.Second) // 每隔两秒返回一个时间
	// for t := range timer {              // 循环定时器
	// 	fmt.Println(t)
	// }

	// 按照指定格式解析时间字符串
	// timeObj, err := time.Parse("2006-01-02 15:04:05", "2020-02-07 13:10:10") // 这个解析出来的时间可能不对，要用下面的方式
	timeObj, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-02-07 10:10:10", time.Local)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	// 打印出上面解析出来的时间的时间戳
	fmt.Println(timeObj.Unix())             // 1581041410
	fmt.Println(TimeFormat(timeObj.Unix())) // 2020-02-07 10:10:10

	fmt.Println("当前时间加两个小时")

	now1 := time.Now()
	fmt.Println(TimeFormat2(now1,""))
	t := now1.Add(time.Hour * time.Duration(2))
	fmt.Println(TimeFormat2(t,""))
	fmt.Println()

	// 两个时间相减
	dur := now.Sub(timeObj)
	fmt.Println(dur.String())  // 2h22m2.7206038s
	fmt.Println(dur.Seconds()) // 8522.7206038 秒数

	// 睡眠 Sleep
	time.Sleep(time.Duration(5 * time.Second)) // Sleep() 接收的参数是纳秒数的 time.Duration 类型，5 * time.Second 是5秒的纳秒数，然后强制转换成 time.Duration 类型（Duration其实是 int64类型）
	fmt.Println("5秒过去了。。。")
}

// TimeFormat 时间格式化
func TimeFormat2(t time.Time, f string) string {
	if len(f) == 0 {
		f = "2006-01-02 15:04:05"
	}
	return t.Format(f)
}