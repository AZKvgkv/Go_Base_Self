package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

// 定义方法:

func (s Student) test01() {
	fmt.Println(s.Name, "test01")
}

//定义函数

func method01(s Student) {
	fmt.Println(s.Name, "method01")
}

func main() {
	//调用函数
	var s = Student{Name: "Tom"}
	method01(s)
	fmt.Println("Method")
	//调用方法
	s.test01()
}
