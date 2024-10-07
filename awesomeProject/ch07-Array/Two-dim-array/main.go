package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	//// 二维数组
	//var arr [3][3]int
	//fmt.Println(arr)
	//
	//// 初始化二维数组
	//arr = [3][3]int{{1, 2}, {3, 4, 5}}
	//arr[0][2] = 9
	//fmt.Println(arr)

	// 二维数组的遍历
	var array = [3][3]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}

	//方式一：普通 for 循环
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			fmt.Print(array[i][j], "\t")
		}
		fmt.Println()
	}
	fmt.Println("--------------------------")
	//方式二：for range 循环
	for i, v1 := range array {
		for j, v2 := range v1 {
			fmt.Printf("array[%v][%v] = [%v]\t", i, j, v2)
		}
		fmt.Println()
	}

}
