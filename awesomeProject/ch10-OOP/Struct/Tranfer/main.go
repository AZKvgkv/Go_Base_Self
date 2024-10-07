package main

import "fmt"

func main() {
	// 结构体的转换
	var s Student = Student{18}
	var p Person = Person{24}
	fmt.Println(s)
	s = Student(p)
	fmt.Println(s)
	fmt.Println(p)
}

type Student struct {
	Age int
}

type Person struct {
	Age int
}
