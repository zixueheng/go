package main

import (
	"encoding/json"
	"fmt"
)

// 在Go语言中，可以使用 json.Marshal() 函数将结构体格式的数据格式化为 JSON 格式。

func main() {
	// 声明技能结构体
	type Skill struct {
		Name string
		// 在转换 JSON 格式时，JSON 的各个字段名称默认使用结构体的名称，如果想要指定为其它的名称我们可以在声明结构体时添加一个`json:" "`标签：
		Level int `json:"Grade"`
	}
	// 声明角色结构体
	type Actor struct {
		id     int
		Name   string
		Age    int
		Skills []Skill // 切片
	}
	// 填充基本角色数据
	a := Actor{
		id:   1, // 字段名需要首字母大写，否则 json 包 取不到字段值（首字母小写相当于私有，大写公有）
		Name: "cow boy",
		Age:  37,
		Skills: []Skill{
			{Name: "Roll and roll", Level: 1},
			{Name: "Flash your dog eye", Level: 2},
			{Name: "Time to have Lunch", Level: 3},
		},
	}
	result, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	jsonStringData := string(result)
	fmt.Println(jsonStringData)
	// {"Name":"cow boy","Age":37,"Skills":[{"Name":"Roll and roll","Grade":1},{"Name":"Flash your dog eye","Grade":2},{"Name":"Time to have Lunch","Grade":3}]}
}

// `json:" "` 标签的使用总结为以下几点：
//     FieldName int `json:"-"`：表示该字段被本包忽略；
//     FieldName int `json:"myName"`：表示该字段在 JSON 里使用“myName”作为键名；
//     FieldName int `json:"myName,omitempty"`：表示该字段在 JSON 里使用“myName”作为键名，并且如果该字段为空时将其省略掉；
//     FieldName int `json:",omitempty"`：该字段在json里的键名使用默认值，但如果该字段为空时会被省略掉，注意 omitempty 前面的逗号不能省略。
