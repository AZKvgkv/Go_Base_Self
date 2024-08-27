package main

import "fmt"

func main() {
	// 算术运算符
	var a = 21
	var b = 12
	var c int

	c = a + b // 加法
	fmt.Printf("第一行 - c 的值是 %d\n", c)
	c = a - b // 减法
	fmt.Printf("第二行 - c 的值是 %d\n", c)
	c = a * b // 乘法
	fmt.Printf("第三行 - c 的值是 %d\n", c)
	c = a / b // 除法
	fmt.Printf("第四行 - c 的值是 %d\n", c)
	c = a % b // 取余
	fmt.Printf("第五行 - c 的值是 %d\n", c)

	a++ // 自增
	fmt.Printf("第六行 - a 的值是 %d\n", a)
	a-- // 自减
	fmt.Printf("第七行 - a 的值是 %d\n", a)

	// 逻辑运算符
	//var x = true
	//var y = false
	//
	//if x && y {
	//	fmt.Printf("第一行 - 条件为 true\n")
	//}
	//if x || y {
	//	fmt.Printf("第二行 - 条件为 true\n")
	//}
	///* 修改 x 和 y 的值 */
	//x = false
	//y = true
	//if x && y {
	//	fmt.Printf("第三行 - 条件为 true\n")
	//} else {
	//	fmt.Printf("第三行 - 条件为 false\n")
	//}
	//if !(x && y) {
	//	fmt.Printf("第四行 - 条件为 true\n")
	//}

	fmt.Println("------逻辑运算符的用法-----")
	var q uint = 60 /* 60 = 0011 1100 */
	var r uint = 13 /* 13 = 0000 1101 */
	var s uint = 0
	s = q & r /* 12 = 0000 1100 */
	fmt.Printf("第一行 - s 的值为 %d\n", s)

	s = q | r /* 61 = 0011 1101 */
	fmt.Printf("第二行 - s 的值为 %d\n", s)

	s = q ^ r /* 49 = 0011 0001 */
	fmt.Printf("第三行 - s 的值为 %d\n", s)

	s = q << 2 /* 240 = 1111 0000 */
	fmt.Printf("第四行 - s 的值为 %d\n", s)

	s = q >> 2 /* 15 = 0000 1111 */
	fmt.Printf("第五行 - s 的值为 %d\n", s)

	fmt.Println("------其他运算符的用法-----")
	var a1 = 10
	var b1 int32
	var c1 float32
	var ptr *int
	/* 运算符实例 */
	fmt.Printf("第 1 行 - a_1 变量类型为 = %T\n", a1)
	fmt.Printf("第 1 行 - b_1 变量类型为 = %T\n", b1)
	fmt.Printf("第 1 行 - c_1 变量类型为 = %T\n", c1)

	/* 指针实例 */
	ptr = &a1
	fmt.Printf("a_1 的值为  %d\n", a1)
	fmt.Printf("*ptr 为 %d\n", *ptr)
}
