package main

import (
	"fmt"
	"reflect"
)

// 利用一个函数，函数的参数定义为空接口
func testReflect(x interface{}) {
	fmt.Println(x)
	// 1. 调用TypeOf()函数，返回x的类型
	reflectType := reflect.TypeOf(x)
	fmt.Println("reType: ", reflectType)
	// 2. 调用ValueOf()函数，返回x的值
	reflectValue := reflect.ValueOf(x)
	fmt.Println("reValue: ", reflectValue)

	// reValue 转换成空接口
	i2 := reflectValue.Interface()
	// 类型断言
	if i2, ok := i2.(Student); ok {
		fmt.Println(i2)
	} else {
		fmt.Println("类型断言失败")
	}
}

type Student struct {
	Name string
	Age  int
}

func main() {
	stu := Student{"Tom", 18}

	testReflect(stu)
}
