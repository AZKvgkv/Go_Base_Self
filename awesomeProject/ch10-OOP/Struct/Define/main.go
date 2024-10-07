package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	// 方式一:按照顺序赋值操作
	var stu1 = Student{Name: "小李", Age: 18}
	fmt.Println(stu1)

	// 方式二:使用键值对赋值(指定类型)
	var stu2 = Student{
		Name: "Tom",
		Age:  20,
	}
	fmt.Println(stu2)

	// 方式三:返回结构体的指针类型
	stu3 := &Student{
		Name: "Jerry",
		Age:  25,
	}
	fmt.Println(*stu3)
}
