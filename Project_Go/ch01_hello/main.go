package main

import "fmt"

// 使用 const 来定义枚举类型(类似Rust语言中的枚举)
const (
	RED   = 1
	GREEN = 2
	BLUE  = 3
)

const (
	BEIJING = 10 * iota
	SHANGHAI
	GUANGZHOU
)

func main() {
	fmt.Println("Hello, world!")
	// 变量声明
	// 方法一 : 声明一个默认值为 0 的变量
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a is %T\n", a)

	// 方法二 : 声明一个变量, 初始化一个值
	var b int = 314
	fmt.Printf("b = %d, type of b is %T\n", b, b)

	var bb string = "hello golang vs_code"
	fmt.Printf("bb = %s, type of bb is %T\n", bb, bb)

	// 方法三 : 初始化的时候可以直接省去变量的类型
	var c = 521
	fmt.Printf("c = %d, type of c is %T\n", c, c)

	var cc = "hello world"
	fmt.Printf("cc = %s, type of cc is %T\n", cc, cc)

	// 方法四 : (常用的方法), 使用海象符
	// 注 : 不支持全局变量
	d := 1_000
	fmt.Printf("d = %d, type of d is %T\n", d, d)

	dd := "hello world go lang"
	fmt.Printf("dd = %s, type of dd is %T\n", dd, dd)

	// 多变量声明
	var xx, yy int = 1, 2
	var ff, kk = 9.8, "number"
	fmt.Printf("xx = %d, yy = %d, ff = %f, kk = %s\n", xx, yy, ff, kk)
	var (
		nn int  = 19
		mm bool = false
	)
	fmt.Print("nn = ", nn, ", mm = ", mm)

	/* 常量的声明 */
	// 常量(只读属性)
	const length int = 10
	fmt.Printf("length = %d, type of length is %T\n", length, length)

	fmt.Println("Red = ", RED)
	fmt.Println("Green = ", GREEN)
	fmt.Println("Blue = ", BLUE)

	fmt.Println("Beijing = ", BEIJING)     // 0
	fmt.Println("Shanghai = ", SHANGHAI)   // 10
	fmt.Println("Guangzhou = ", GUANGZHOU) //20
}
