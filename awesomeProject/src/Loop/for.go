package main

import "fmt"

func main() {
	/*
		// 实现对于10以内数字求和
		sum := 0
		for i := 1; i <= 10; i++ {
			sum += i
		}
		println(sum)
	*/

	// 实现一个功能：输出1~100之内能够被7整除的数
	//for i := 1; i <= 100; i++ {
	//	if i % 7 == 0 {
	//		println(i)
	//	}
	//}

	//实现一个功能：输出1~100之内能够被6整除的数
	for i := 1; i <= 100; i++ {
		if i%6 != 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}
