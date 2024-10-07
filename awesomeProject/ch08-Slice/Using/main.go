package main

import "fmt"

func main() {
	// 数组的遍历
	// 方式一： for
	// 方式二： for range

	// 定义切片
	slice := make([]int, 4, 20)
	slice[0] = 66
	slice[1] = 77
	slice[2] = 88
	slice[3] = 99
	// 方式1 普通for循环
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v \n", i, slice[i])
	}
	// 方式2 for range
	for index, value := range slice {
		fmt.Println(index, value)
	}

}
