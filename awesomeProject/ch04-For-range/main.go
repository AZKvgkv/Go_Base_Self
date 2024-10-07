package main

import "fmt"

func main() {
	// 定义一个字符串
	var str string = "halo golang"
	// 方式一：普通for循环,按照字节进行遍历输出
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}
	fmt.Println("-------------------------")
	// 方式二：range
	for i, v := range str {
		fmt.Printf("%c", v)
		fmt.Printf("索引为：%d, 具体的值为：%c\n", i, v)
	}
}
