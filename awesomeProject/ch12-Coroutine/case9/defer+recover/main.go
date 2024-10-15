package main

import (
	"fmt"
	"time"
)

func main() {
	go printNum(3)
	go div(10, 2)
	go div(10, 0)
	time.Sleep(time.Second)
}

// 输出数字
func printNum(num int) {
	for i := 1; i <= num; i++ {
		fmt.Println(i)
	}
}

// 做除法操作
func div(a, b int) {
	func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println(a / b)
}
