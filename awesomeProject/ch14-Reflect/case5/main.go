package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) A_Print() {
	fmt.Println(s.Name, s.Age)
}

func (s Student) B_GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s *Student) C_Set(name string, age int) {
	s.Name = name
	s.Age = age
}
func TestStudentStruct(a interface{}) {
	reflectValue := reflect.ValueOf(a)
	fmt.Println(reflectValue)

	// 通过reflect.Value类型操作结构体内部的字段
	n1 := reflectValue.NumField()
	fmt.Println(n1)
	for i := 0; i < n1; i++ {
		fmt.Println(reflectValue.Field(i))
	}

	// 通过reflect.Value类型操作结构体内部的方法
	n2 := reflectValue.NumMethod()
	fmt.Println(n2)

	fmt.Println("-------1--------")
	reflectValue.Method(0).Call(nil)

	fmt.Println("-------2--------")
	params := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	// var params []reflect.Value
	// params = append(params, reflect.ValueOf(10))
	// params = append(params, reflect.ValueOf(20))
	result := reflectValue.Method(1).Call(params)
	fmt.Println(result[0].Int())
}
func main() {
	s := Student{"Tom", 18}
	TestStudentStruct(s)
}
