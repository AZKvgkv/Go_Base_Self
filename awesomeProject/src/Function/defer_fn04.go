package main

import "fmt"

func main() {
	add(3, 66)
}

// 在Go语言中，defer语句会延迟函数的执行直到退出函数，在defer语句执行时，函数的参数已经求值完毕。
// 程序遇到defer语句时，会先将defer语句压入栈中，当函数执行到最后时，会从栈中依次取出defer语句执行。(先进后出)

func add(a, b int) int {
	defer fmt.Println("a = ", a)
	defer fmt.Println("b = ", b)
	var sum = a + b
	fmt.Println("sum = ", sum)
	return sum
}
