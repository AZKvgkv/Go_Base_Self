package main

import "fmt"
import "awesomeProject/ch10-OOP/Struct-Ins-package/model"

func main() {
	// 跨包创建结构体Student的实例：
	s := model.Student{Name: "Az", Age: 24}
	fmt.Println(s)

	s1 := model.NewStudent("YX", 25)
	fmt.Println(*s1)
}
