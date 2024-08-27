package main

import "fmt"

func main() {
	// 双重循环
label1:
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			fmt.Printf("%d * %d = %d\n", i, j, i*j)
			if i == 2 && j == 2 {
				break label1 //结束指定标签的循环
			}
		}
	}
	fmt.Printf("--------ok--------")
}
