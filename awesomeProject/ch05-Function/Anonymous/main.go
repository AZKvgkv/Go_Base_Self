package main

import "fmt"

func main() {
	//匿名函数:定义的同时调用(只希望使用一次！)
	f := func(x, y int) int {
		return x + y
	}(11, 26)
	fmt.Printf("%v\n", f)
}