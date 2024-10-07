package main

import "fmt"

func main() {
	fmt.Println("变量的声明")
	var (
		num1 int
		num  = 18
		num2 = "tom" //也可以声明变量的类型
	)
	sex := "男"
	fmt.Println(num1, num, num2, sex)

}
