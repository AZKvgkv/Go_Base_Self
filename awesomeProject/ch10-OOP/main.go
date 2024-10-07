package main

import "fmt"

func main() {
	// 面向对象编程
	// Go 语言支持面向对象编程（OOP）语法，Go 语言没有类。使用结构体（struct）和方法（method）实现面向对象编程。

	// 创建老师结构体的实例、对象、变量：
	var T1 Teacher
	T1.Name = "马士兵"
	T1.Age = 45
	T1.School = "清华大学"

	var T2 Teacher
	T2.Name = "赵珊珊"
	T2.Age = 31
	T2.School = "黑龙江大学"

	fmt.Println(T1, T2)
}

// Teacher 定义老师结构体，将老师中的各个属性  统一放到一个结构体中管理：
type Teacher struct {
	Name   string
	Age    int
	School string
}
