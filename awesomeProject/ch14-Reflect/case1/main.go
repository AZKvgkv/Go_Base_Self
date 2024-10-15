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

	num := reflectValue.Int()
	fmt.Println(num * 2)

	// reValue 转换成空接口
	i2 := reflectValue.Interface()
	n := i2.(int)
	fmt.Println(n)

}

func main() {
	var num int = 42
	testReflect(num)
}
