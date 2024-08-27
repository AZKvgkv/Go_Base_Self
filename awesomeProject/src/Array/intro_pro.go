package main

import "fmt"

func main() {
	// 实现的功能：将给出的5个学生的成绩，计算其总和以及平均值
	// 使用数组遍历的方式
	var scores [5]int
	// 输入5个学生的成绩，将成绩存入数组
	for i := 0; i < len(scores); i++ {
		fmt.Printf("请输入第%d个学生的成绩：\n", i+1)
		_, err := fmt.Scanln(&scores[i])
		if err != nil {
			return
		}
	}

	//sum := 0
	//for i := 0; i < len(scores); i++ {
	//	sum += scores[i]
	//}
	//avg := float64(sum) / float64(len(scores))
	//
	//fmt.Printf("成绩总和为：%d\n成绩平均值为：%.2f\n", sum, avg)

	// 输出每个学生的成绩
	// 方式一：普通for循环
	for i := 0; i < len(scores); i++ {
		fmt.Printf("第%d个学生的成绩为：%d\n", i+1, scores[i])
	}

	fmt.Println("----------------------------")

	// 方式二：for-range循环
	for index, value := range scores {
		fmt.Printf("第%d个学生的成绩为：%d\n", index+1, value)
	}

}
