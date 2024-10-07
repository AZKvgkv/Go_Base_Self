package main

import "fmt"

func main() {
	test()
	fmt.Println("程序继续执行")
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("程序出现了:%v\n", err)
		}
	}()
	num1 := 10
	num2 := 5             //如果改成0就报错了
	result := num1 / num2 // 结果为2，因为整数除法只保留整数部分
	fmt.Println("结果是:", result)
}
