package model

import "sync"

// 1) 懒汉式，需要的时候初始化实例

// Tool 结构体模拟类
type Tool struct{}

// String 返回描述字符串
func (t *Tool) String() string {
	return "I am a tool"
}

// 私有变量
var instance *Tool

// lock 锁对象，进行加锁保证线程安全，但由于每次调用该方法都进行了加锁操作，在性能上不是很高效。
var lock sync.Mutex

// once 对象 可以确保操作只进行一次
var once sync.Once

// GetInstance 获取实例
// func GetInstance() *Tool {
// 	lock.Lock() // 加锁，保证线程安全
// 	defer lock.Unlock()
// 	if instance == nil {
// 		return new(Tool)
// 	}
// 	return instance
// }

// GetInstance 获取实例（双重检查、改进款）
// 第一次判断不加锁，第二次加锁保证线程安全，一旦对象建立后，获取对象就不用加锁了
// 在懒汉式（线程安全）的基础上再进行优化，减少加锁的操作，保证线程安全的同时不影响性能。
func GetInstance() *Tool {
	if instance == nil {
		lock.Lock()
		if instance == nil {
			return new(Tool)
		}
		lock.Unlock()
	}
	return instance
}

// GetInstanceOnce 通过 sync.Once 来确保创建对象的方法只执行一次
// sync.Once 内部本质上也是双重检查的方式，只是用来替代上面的书写方式
func GetInstanceOnce() *Tool {
	once.Do(func() {
		instance = new(Tool)
	})
	return instance
}
