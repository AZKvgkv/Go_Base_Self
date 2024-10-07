package main

import "fmt"

func main() {
	//  闭包:返回的匿名函数 + 匿名函数以外的变量（比如这里的 num）
	//  闭包的作用：可以读取函数内部的变量
	//  闭包的特性：匿名函数引用了外部函数的变量，所以匿名函数的变量是引用类型，所以闭包可以读取函数内部的变量
	//  匿名函数中的变量，在函数执行结束后，函数执行的内存空间会被释放，但是匿名函数引用的变量，因为被匿名函数引用，所以不会被释放
	f := getSum()
	fmt.Println(f(1))
	fmt.Println(f(41))
}
func getSum() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}
