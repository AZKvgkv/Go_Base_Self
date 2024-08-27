package main

import "fmt"

func main() {
	sum := cal(56, 78)
	fmt.Printf("sum = %d\n", sum)
}
func cal(num1 int, num2 int) int {
	//如果返回值类型就一个的话，那么第二个括号可以省略
	var sum = 0
	sum = num1 + num2
	return sum
}
