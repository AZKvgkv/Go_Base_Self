package main

import "fmt"

// 定义动物结构体
type Animal struct {
	Age    int
	Weight float32
}

// 给动物结构体添加方法 喊叫
func (an *Animal) Shout() {
	println("动物叫")
}
// 给动物结构体添加方法 自我展示
func (an *Animal) ShowInfo() {
	fmt.Printf("我是一只动物，我的年龄是%v，我的体重是%v\n",an.Age, an.Weight)
}

// 定义猫结构体
type Cat struct {
	Animal // 继承Animal结构体
}

// 给猫结构体添加方法 跑
func (c *Cat) Run() {
	println("猫跑")
}

// 给猫结构体添加方法 吃鱼
func (c *Cat) EatFish() {
	println("猫吃鱼")
}


func main() {
	// 实例化一个猫
	// cat := Cat{Animal{2, 3.5}, "橘色"}
	// // 调用猫的方法
	// cat.Shout()
	// cat.ShowInfo()
	// cat.Run()
	// cat.EatFish()
	cat := &Cat{}
	cat.Animal.Age = 2
	cat.Animal.Weight = 3.5
	cat.Animal.Shout()
	cat.Animal.ShowInfo()
	cat.Run()
	cat.EatFish()
}