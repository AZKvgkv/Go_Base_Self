package main

import "fmt"

//  构造函数先于主函数被执行
//  step 01 --> 全局变量
//  step 02 --> init函数
//  step 03 --> main函数

func main() {
	fmt.Println("main函数被执行！")

	var number = test()
	fmt.Println("number:", number)
}

func init() {
	fmt.Println("init函数被执行！")
}

func test() int {
	fmt.Println("test函数被执行！")
	return 666
}
