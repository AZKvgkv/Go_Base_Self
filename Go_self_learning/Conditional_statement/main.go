package main

import "fmt"

func main() {
	var grade = "B"
	var marks = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)

	/* for 循环 */
	var x = 15
	var y int
	numbers := [6]int{1, 2, 3, 5}

	for y := 0; y < 10; y++ {
		fmt.Printf("y 的值为:%d\n", y)
	}
	for y < x{
		y++
		fmt.Printf("y 的值为:%d\n", y)
	}
	for i,n := range numbers{
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,n)
	}
}
