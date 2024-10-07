package main

import "fmt"

func main() {
	// map的操作： 增删改查
	// 增加
	// map[key] = value

	// 删除
	// delete(map, key)

	// 修改
	// map[key] = value

	// 查找
	// value, ok := map[key]

	// 实际代码展示
	// 增加
	var map1 map[string]int
	map1 = make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3

	// 修改
	map1["one"] = 11
	fmt.Println(map1)

	// 删除
	delete(map1, "one")
	fmt.Println(map1)

	// 查找
	value, ok := map1["two"]
	if ok {
		fmt.Printf("The %v corresponding to key_two\n", value)
	} else {
		fmt.Println("key not found")
	}

	// 清空
	map1 = make(map[string]int)
	fmt.Println(map1)

	// 调用函数
	mapPro()
}

// map的进一步操作
func mapPro() {
	// 新建一个map
	e := map[int]string{
		1: "AZ",
		2: "YX",
		3: "WJ",
	}
	fmt.Println(e)
	// 获取长度
	fmt.Println(len(e))

	// 使用for - range遍历map
	for k, v := range e {
		fmt.Println(k, v)
	}

	// pro 版本
	f := make(map[string]map[int]string)

	// 赋值
	f["one"] = make(map[int]string, 3)
	f["one"][1] = "AZ"
	f["one"][2] = "YX"
	f["one"][3] = "WJ"
	f["two"] = make(map[int]string, 3)
	f["two"][1] = "CXQ"
	f["two"][2] = "SKX"
	f["two"][3] = "ZU"

	for k1, v1 := range f {
		fmt.Println(k1)
		for k2, v2 := range v1 {
			fmt.Printf("学生的序号为：%v 学生的姓名为：%v \t", k2, v2)
		}
		fmt.Println()
	}
}
