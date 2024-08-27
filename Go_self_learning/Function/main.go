package main

import (
	"fmt"
	"math"
)

func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func swap(x, y string) (string, string) {
	return y, x
}

// 方法
/* 定义函数 */
type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	var a = 12
	var b = 13
	var c = max(a, b)
	fmt.Println("The maximum of", a, "and", b, "is", c)

	str1, str2 := swap("hello", "golang")
	fmt.Println(str1, str2)

	// 函数作为值
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))

	// 闭包
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 的值将会递增 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())

	var c1 Circle
	c1.radius = 100
	fmt.Println("Area of Circle(c1) = ", c1.getArea())


	// defer 的使用
	fmt.Print("Start\n")
	defer fmt.Print("1\n")
	defer fmt.Print("2\n")
	defer fmt.Print("3\n")
	fmt.Print("End\n")

}
