package main

import "fmt"

func main() {
	add(1, 2)
}

func add(a, b int) int {
	defer fmt.Println("a = ", a)
	defer fmt.Println("b = ", b)
	var sum = a + b
	fmt.Println("sum = ", sum)
	return sum
}
