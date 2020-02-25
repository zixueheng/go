package main

import "fmt"

// 实例化一个字符串映射的 函数切片
// 创建一个 map 实例，这个 map 通过事件名（string）关联回调列表（[]func(interface{})），同一个事件名称可能存在多个事件回调，因此使用回调列表保存。回调的函数声明为 func(interface{})。
var eventByName = make(map[string][]func(interface{}))

// 1、事件注册
// 事件系统需要为外部提供一个注册入口。这个注册入口传入注册的事件名称和对应事件名称的响应函数，事件注册的过程就是将事件名称和响应函数关联并保存起来

// eventRegister 事件注册函数
// name 事件名
// callback 回调函数
func eventRegister(name string, callback func(interface{})) {
	list := eventByName[name]     // 取出这个name的 函数切片
	list = append(list, callback) // 将callback 追加的 list 切片中
	eventByName[name] = list      // 再将更新过的 list 切片 保存回去
}

// 2、事件调用
// 事件调用方和注册方是事件处理中完全不同的两个角色。
// 事件调用方是事发现场，负责将事件和事件发生的参数通过事件系统派发出去，而不关心事件到底由谁处理；事件注册方通过事件系统注册应该响应哪些事件及如何使用回调函数处理这些事件

// callEvent 调用事件
// name 事件名
// param 参数，描述事件具体的细节，例如门打开的事件触发时，参数可以传入谁进来了。
func callEvent(name string, param interface{}) {
	list := eventByName[name] // 通过名字查找事件列表
	// 遍历这个事件的所有回调
	for _, callback := range list {
		// 传入参数调用回调
		callback(param)
	}
}

// 3、使用事件

// actor 角色结构体
type actor struct{}

// onEvent actor的事件处理方法，param 参数，类型为 interface{}，与事件系统的函数（func(interface{})）签名一致
func (a *actor) onEvent(param interface{}) {
	fmt.Println("actor event:", param)
}

// gobalEvent 全局事件，有时需要全局进行侦听或者处理一些事件，这里使用普通函数实现全局事件的处理
func gobalEvent(param interface{}) {
	fmt.Println("global event:", param)
}

func main() {
	// 实例化一个角色
	a := new(actor)
	// 注册名为OnSkill的回调，实现代码由 a 的 onEvent 进行处理。也就是 Actor的OnEvent() 方法。
	eventRegister("onSkill", a.onEvent)
	// 再次在OnSkill上注册全局事件，实现代码由 globalEvent 进行处理，虽然注册的是同一个名字的事件，但前面注册的事件不会被覆盖，而是被添加到事件系统中，关联 OnSkill 事件的函数列表中。
	eventRegister("onSkill", gobalEvent)
	// 调用事件，所有注册的同名函数都会被调用
	callEvent("onSkill", 100)
	// actor event: 100
	// global event: 100
}
