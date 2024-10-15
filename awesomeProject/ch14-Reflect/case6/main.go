package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func TestStudentStruct(a interface{}) {
	reflectValue := reflect.ValueOf(a)
	fmt.Println(reflectValue)

	reflectValue.Elem().Field(0).SetString("Jerry")

}
func main() {
	s := Student{"Tom", 18}
	TestStudentStruct(&s)
	fmt.Println(s)
}
