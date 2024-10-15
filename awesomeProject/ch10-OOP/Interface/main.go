package main

import "fmt"

// 接口的定义：定义规则、定义规范、定义某种能力

type SayHello interface {
	// 声明没有实现的方法
	sayHello()
}

// 接口的实现：定义一个结构体
// 中国人

type Chinese struct {
}

// 实现接口的方法 --> 具体的实现
func (c *Chinese) sayHello() {
	println("你好，世界！")
}

// 英国人
type English struct {
}

func (e *English) sayHello() {
	println("Hello, world!")
}

// 定义一个函数：专门用来调用接口的sayHello方法
func greet(s SayHello) {
	s.sayHello()
}

// 自定义数据类型
type integer int

func (i integer) sayHello() {
	fmt.Printf("say %d Hello!\n", i)
}

func main() {
	var i integer = 10
	var say SayHello = i
	say.sayHello()
	fmt.Println("----------")
	greet(&Chinese{})
	greet(&English{})
	c := Chinese{}
	var s SayHello = &c
	s.sayHello()
}
