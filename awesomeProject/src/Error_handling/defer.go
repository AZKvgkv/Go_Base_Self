package main

import "fmt"

func main() {
	test()
	fmt.Println("程序继续执行")
}

func test() {
	defer func() {
		// 调用recover函数，可以捕获到panic抛出的错误信息
		if err := recover(); err != nil {
			fmt.Println("发生错误:", err)
		}
	}()
	num1 := 10
	num2 := 5             //如果改成0就报错了
	result := num1 / num2 // 结果为2，因为整数除法只保留整数部分
	fmt.Println("结果是:", result)
}
