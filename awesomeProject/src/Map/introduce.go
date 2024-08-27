package main

import "fmt"

func main() {
	// map, key-value, 类似python中的字典
	// 基本语法：var mapName map[keyType]valueType
	// map的特点：（1） 使用前需要make，（2） map中的key是无序的，（3） map中的key是唯一的，（4） map中的value可以重复

	var a map[int]string

	a = make(map[int]string, 10) // 10 为容量——可以放10个键值对
	a[1] = "AZ"
	a[2] = "YX"
	a[3] = "WJ"

	fmt.Println(a)

	// map 的三种创建方式
	// 1. make
	b := make(map[int]string, 10)
	b[1] = "AZ"
	b[2] = "YX"
	b[3] = "WJ"
	fmt.Println(b)

	// 2. 直接赋值
	var c = map[int]string{
		1: "AZ",
		2: "YX",
		3: "WJ",
	}
	fmt.Println(c)

	// 3. 简写
	d := map[int]string{
		1: "AZ",
		2: "YX",
		3: "WJ",
	}
	fmt.Println(d)
}
