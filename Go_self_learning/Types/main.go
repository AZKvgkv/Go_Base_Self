package main

import "fmt"

func main() {
	fmt.Println("---Slice part---")
	slice_part()

	slice_bound()

	slice_append_copy()

	range_func()

	map_func()

	delete_map_element()
}

func slice_part() {
	numbers := make([]int, 3, 5)
	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(x), cap(x), x)
}

func slice_bound() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	// 打印原始切片
	fmt.Println("numbers == ", numbers)

	// 截取索引为1~4的切片
	fmt.Println("numbers[1:4] == ", numbers[1:4])
}

func slice_append_copy() {
	var numbers []int
	printSlice(numbers)

	/*允许追加空切片*/
	numbers = append(numbers, 0)
	printSlice(numbers)

	/*追加元素到切片*/
	numbers = append(numbers, 1, 2, 3)
	printSlice(numbers)

	/*创建容量是之前两倍的切片*/
	numbers_new := make([]int, len(numbers), (cap(numbers))*2)

	// 拷贝旧切片的内容到新的切片
	copy(numbers_new, numbers)
	printSlice(numbers_new)
}

func range_func() {
	fmt.Println("---Range part---")
	nums := []int{1, 6, 8}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum = ", sum)

	// 需要知道索引
	for index, num := range nums {
		fmt.Println("index = ", index, "num = ", num)
	}

	// range也可以使用在map的键值对上
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range也可以用来枚举Unicode字符串，返回字符和它的索引
	for index, ch := range "Go语言" {
		fmt.Printf("%d -> %c\n", index, ch)
	}
}

func map_func() {
	// 创建集合
	countryCapitalMap := make(map[string]string)

	// map插入key - value对,各个国家对应的首都
	countryCapitalMap["China"] = "Beijing"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"
	countryCapitalMap["USA"] = "Washington"
	countryCapitalMap["Russia"] = "Moscow"
	countryCapitalMap["Germany"] = "Berlin"
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Brazil"] = "Brasilia"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Spain"] = "Madrid"
	countryCapitalMap["Canada"] = "Ottawa"
	countryCapitalMap["South Korea"] = "Seoul"
	countryCapitalMap["Australia"] = "Canberra"

	// 使用key输出map值
	for country := range countryCapitalMap {
		fmt.Println("Capital of ", country, "is", countryCapitalMap[country])
	}

	// 查看元素在集合中是否存在
	capital, ok := countryCapitalMap["United States"]
	if ok {
		fmt.Println("Capital of ", "United States", "is", capital)
	} else {
		fmt.Println("Country United States not found")
	}
}

func delete_map_element() {
	fmt.Println("---Delete map element---")
	// 创建 map
	countryCapitalMap := map[string]string{
		"France":      "Paris",
		"Brazil":      "Brasilia",
		"Italy":       "Rome",
		"Spain":       "Madrid",
		"Canada":      "Ottawa",
		"South Korea": "Seoul",
		"Japan":       "Tokyo",
		"Australia":   "Canberra",
	}
	fmt.Println("原始 map")

	// 打印原始 map
	for country := range countryCapitalMap {
		fmt.Println("Capital of ", country, "is", countryCapitalMap[country])
	}

	// delete 元素
	delete(countryCapitalMap, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后 map")

	// 打印原始 map
	for country := range countryCapitalMap {
		fmt.Println("Capital of ", country, "is", countryCapitalMap[country])
	}

}
