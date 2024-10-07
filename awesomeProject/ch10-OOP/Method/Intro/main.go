package main

import "fmt"

func main() {
	// 创建结构体对象
	//var p Person = Person{"Tom"}
	p := Person{"Tom"}
	p.test()
}

type Person struct {
	Name string
}

func (p Person) test() {
	println("Person test")
	fmt.Println(p.Name)
}
