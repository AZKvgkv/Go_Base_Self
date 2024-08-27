package main

import (
	"fmt"
	"unsafe"
)

// 变量的声明
var a = "it is from w3c_school"
var b = "这是一个字符串 & w3c_school"
var c bool

// 多变量声明
var x, y int
var (
	_a int
	_b bool
)
var _c, d = 1, 2
var e, f = 123, "hello"

// 常量还可以用作枚举
const (
	A = "abc"
	B = len(A)
	C = unsafe.Sizeof(A)
)

func main() {
	println(a, b, c)
	g, h := 100, "world"
	println(x, y, _a, _b, _c, d, e, f, g, h)

	// 常量
	const LENGTH int = 10
	const (
		WIDTH  = 100
		HEIGHT = 50
	)
	var area = LENGTH * WIDTH
	fmt.Printf("面积为：%d\n", area)
	var perimeter = 2 * (LENGTH + WIDTH)
	fmt.Printf("周长为：%d\n", perimeter)
	var volume = LENGTH * WIDTH * HEIGHT
	fmt.Printf("体积为：%d\n", volume)

	println("枚举示例：", A, B, C)

	// iota 的用法
	const (
		I = iota
		KB
		MB
		GB = "ha"
		TB
		PB = 100
		EB
		ZB = iota
		YB
	)
	fmt.Println("value of all:", I, KB, MB, GB, TB, PB, EB, ZB, YB)

	// iota 实例
	const (
		i = 1 << iota // 这里是二进制左移 注: << n == *(2^n)
		j = 3 << iota
		k
		l
	)

	fmt.Println("value of i and j and k and l:", i, j, k, l)
}
