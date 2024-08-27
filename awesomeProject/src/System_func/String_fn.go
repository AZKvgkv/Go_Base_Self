package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// 统计字符串的长度，按字节进行统计，在golang中，一个汉字是3个字节
	//l := len("hello world")
	//println(l)

	// 对字符串进行遍历
	// 方式一：利用键值循环
	//for i, value := range "hello world" {
	//	println(i, value)
	//}
	//
	//// 方式二：利用for循环
	//for i := 0; i < len("hello world"); i++ {
	//	fmt.Printf("%c\n", string("hello world")[i])
	//}
	//
	//// 方式三：利用r:=[]rune(str)
	//r := []rune("hello world")
	//for i := 0; i < len(r); i++ {
	//	fmt.Printf("%c", r[i])
	//}

	// 字符串转整数
	num1, _ := strconv.Atoi("123")
	fmt.Println(num1)
	// 整数转字符串
	str2 := strconv.Itoa(99)
	fmt.Println(str2)

	// 统计一个字符串有几个指定的子串
	count := strings.Count("hello world", "o")
	fmt.Println(count)
	// 不区分大小写的字符串比较：
	flag :=
		strings.EqualFold("hello", "Hello")
	fmt.Println(flag)

	// 返回子串在字符串第一次出现的索引值，如果没有返回-1：
	index := strings.Index("hello world", "o")
	fmt.Println(index)
	// 返回子串在字符串最后一次出现的索引值，如果没有返回-1：
	index = strings.LastIndex("hello world", "o")
	fmt.Println(index)
}
