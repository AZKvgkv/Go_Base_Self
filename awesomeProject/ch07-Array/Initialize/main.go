package main

import "fmt"

func main() {
	// 数组初始化的方式
	// 1. 指定数组的长度和元素类型
	var arr1 [5]int
	fmt.Printf("arr1的值是：%v\n", arr1)
	// 2. 使用初始化列表进行初始化
	arr2 := [3]int{1, 2, 3}
	fmt.Printf("arr2的值是：%v\n", arr2)
	// 3. 使用索引和值进行初始化
	arr3 := [...]int{4, 5, 6}
	fmt.Printf("arr3的值是：%v\n", arr3)
	// 4. 使用数组字面量进行初始化
	arr4 := [5]int{9, 8, 7}
	fmt.Printf("arr4的值是：%v\n", arr4)
	// 5. 使用指定长度和默认值的数组进行初始化
	arr5 := [5]int{2: 3, 4: 5}
	fmt.Printf("arr5的值是：%v\n", arr5)

}
