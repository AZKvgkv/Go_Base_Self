package main

import (
	"awesomeProject/ch10-OOP/Encapsulation/model"
	"fmt"
)

func main() {
	// 创建person结构体实例
	p := model.NewPerson("Tom")
	p.SetAge(19)
	fmt.Println(p.Name, p.GetAge())
	fmt.Println(*p)
}
