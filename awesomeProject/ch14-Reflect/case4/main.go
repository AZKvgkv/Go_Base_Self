package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 42
	testReflect(&num)
	fmt.Println(num)
}

func testReflect(i interface{}) {
	reValue := reflect.ValueOf(i)
	reValue.Elem().SetInt(12)
}
