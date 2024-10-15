package main

import (
	"fmt"
	"reflect"
)

func main() {
	stu := Student{"Az", 24}
	testReflect(stu)
}
func testReflect(i interface{}) {
	reType := reflect.TypeOf(i)
	reValue := reflect.ValueOf(i)
	// 获取变量的类别
	// way 1
	kind := reType.Kind()
	fmt.Println("reType.Kind: ", kind)
	// way 2
	kind2 := reValue.Kind()
	fmt.Println("reValue.Kind: ", kind2)

	// 获取变量的类型
	i2 := reValue.Interface()
	// 类型断言
	if i2, ok := i2.(Student); ok {
		fmt.Printf("i2结构体的类型是: %T\n", i2)
	} else {
		fmt.Println("类型断言失败")
	}

}

type Student struct {
	Name string
	Age  int
}
