package main

import (
	"encoding/json"
	"fmt"
)

// Screen 定义手机屏幕
type Screen struct {
	Size       float32 // 屏幕尺寸
	ResX, ResY int     // 屏幕水平和垂直分辨率
}

// Battery 定义电池
type Battery struct {
	Capacity int // 容量
}

// genJSONData 生成json数据
func genJSONData() []byte {
	// 完整数据结构
	// 定义了一个匿名结构体。这个结构体内嵌了 Screen 和 Battery 结构体，同时临时加入了 HasTouchID 字段。
	raw := &struct {
		Screen
		Battery
		HasTouchID bool // 序列化时添加的字段：是否有指纹识别
	}{
		// 屏幕参数
		Screen: Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		// 电池参数
		Battery: Battery{
			2910,
		},
		// 是否有指纹识别
		HasTouchID: true,
	}
	// 将数据序列化为json
	jsonData, _ := json.Marshal(raw)
	return jsonData
}

func main() {
	// 生成一段json数据
	jsonData := genJSONData() // 获得 []byte 类型的 JSON 数据
	fmt.Println(string(jsonData))
	// {"Size":5.5,"ResX":1920,"ResY":1080,"Capacity":2910,"HasTouchID":true}

	// 只需要屏幕和指纹识别信息的结构和实例
	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{}
	// 反序列化到screenAndTouch
	json.Unmarshal(jsonData, &screenAndTouch)
	// 输出screenAndTouch的详细结构
	fmt.Printf("%+v\n", screenAndTouch)
	// {Screen:{Size:5.5 ResX:1920 ResY:1080} HasTouchID:true}

	// 只需要电池和指纹识别信息的结构和实例
	batteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}
	// 反序列化到batteryAndTouch
	json.Unmarshal(jsonData, &batteryAndTouch)
	// 输出screenAndTouch的详细结构
	fmt.Printf("%+v\n", batteryAndTouch)
	// {Battery:{Capacity:2910} HasTouchID:true}
}
