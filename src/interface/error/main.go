package main

import (
	"errors"
	"fmt"
	"math"
)

// error 类型是一个非常简单的接口类型，如下所示：
//     // The error built-in interface type is the conventional interface for
//     // representing an error condition, with the nil value representing no error.
//     type error interface {
//         Error() string
//     }

// 一般情况下，如果函数需要返回错误，就将 error 作为多个返回值中的最后一个（但这并非是强制要求）

// Sqrt 平方根
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return -1, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}

func main() {
	result, err := Sqrt(-13)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	// math: square root of negative number

	result2, err2 := Sqrt2(-13)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(result2)
	}
	// Wrong!!!,because "-13.000000" is a negative number
}

// 自定义错误类型
// 即要实现 error 接口中 Error() 方法

// dualError 自定义错误
type dualError struct {
	Num     float64
	problem string
}

// Error 实现 error 接口中的  Error() 方法
func (e dualError) Error() string {
	return fmt.Sprintf("Wrong!!!,because \"%f\" is a negative number", e.Num)
}

// Sqrt2 平方根
func Sqrt2(f float64) (float64, error) {
	if f < 0 {
		return -1, dualError{Num: f}
	}
	return math.Sqrt(f), nil
}
