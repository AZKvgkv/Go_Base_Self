package main

import "fmt"

// 函数名的规范：
// 1. 驼峰命名法
// 2. 首字母不能是数字
// 3. 首字母大写，代表public， 首字母小写，代表private

func main() {
	sum := cal(56, 78)
	fmt.Printf("sum = %d\n", sum)
}

func cal(num1, num2 int) int {
	//如果返回值类型就一个的话，那么第二个括号可以省略
	return num1 + num2
}
