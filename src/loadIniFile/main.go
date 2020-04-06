package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

// 加载本地 ini 配置文件示例

// Mysql 配置
type Mysql struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

// Redis 配置
type Redis struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

// Log 配置
type Log struct {
	Src string
}

// Config 配置内嵌其他配置
type Config struct {
	Mysql `ini:"mysql"`
	Redis `ini:"redis"`
	Log   `ini:"log"`
}

// LoadIni 加载配置文件
// filename 文件名
// data 结构体指针
func LoadIni(filename string, config interface{}) error {
	t := reflect.TypeOf(config)
	if t.Kind() != reflect.Ptr {
		return errors.New("config应该是指针")
	}
	if t.Elem().Kind() != reflect.Struct {
		return errors.New("config应该是指针结构体")
	}
	data, err := ioutil.ReadFile(filename) // 读取文件内容
	if err != nil {
		return err
	}

	lineSlice := strings.Split(string(data), "\r\n") // 以换行拆分成字符串切片
	var configName string                            // 配置命 对应 MySQL Redis Log
	for index, line := range lineSlice {
		// fmt.Println(line)
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") { // 遇到注释跳过
			continue
		}
		if strings.HasPrefix(line, "[") { // 节点名行 []
			if line[0] != '[' || line[len(line)-1] != ']' {
				return fmt.Errorf("第%d行语法错误", index)
			}
			sectionName := strings.TrimSpace(line[1 : len(line)-1]) // 配置文件节点名 [mysql] [redis]

			// 遍历 config 结构体的字段 判断是否和 文件中的节点一致
			// 一致则意味着 该结构体字段 是需要处理的（configName赋值节点名）并在下一次循环的 else 块处理赋值
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					configName = field.Name
				}
			}
		} else { // 非节点名的行
			// TODO
		}
	}
	return nil
}

func main() {
	config := new(Config)
	err := LoadIni("./config.ini", config)
	if err != nil {
		fmt.Println(err)
	}
}
