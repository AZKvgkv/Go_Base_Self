package main

import "fmt"

// func main() {
// 	// 创建结构体对象
// 	//var p Person = Person{"Tom"}
// 	p := Person{"Tom"}
// 	p.test()
// }

// Person 定义Person结构体
type Person struct {
	Name string
}

// 给Person结构体定义一个方法test
func (p Person) test() {
	println("Person test")
	fmt.Println(p.Name)
}
