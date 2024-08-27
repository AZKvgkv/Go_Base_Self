package main

import "fmt"

func main() {
	i := 4
	fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(i))

	for index := 0; index < 10; index++ {
		fmt.Printf("斐波那契数列的第 %d 项是 %d\n", index+1, fibonacci(index))
	}

	type_conversion()
}

func Factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * Factorial(x-1)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func type_conversion() {
	var sum int = 17
	var count int = 5

	mean := float32(sum) / float32(count)
	fmt.Printf("mean = %f\n", mean)
}

