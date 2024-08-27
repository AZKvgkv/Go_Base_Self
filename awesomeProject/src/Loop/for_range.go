package main

import "fmt"

func main() {
	//// 定义一个字符串
	//var str string = "halo golang"
	//// 方式一：普通for循环
	//for i := 0; i < len(str); i++ {
	//	fmt.Printf("%c", str[i])
	//}
	//fmt.Println("-------------------------")
	//// 方式二：range
	//for i, v := range str {
	//	fmt.Printf("%c", v)
	//	fmt.Printf("索引为：%d, 具体的值为：%c\n", i, v)
	//}

	// 实现一个功能：求1~100的和，当和第一次超过300的时候停止程序，输出和已超过300
	var sum = 0
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 300 {
			fmt.Println("和已超过300")
			break
		}
		fmt.Println(sum)
	}
}
