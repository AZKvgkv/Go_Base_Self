package main

import "fmt"

func main() {
	//实现的功能：将给出的5个学生的成绩，计算其总和以及平均值
	score1 := 95
	score2 := 91
	score3 := 63
	score4 := 86
	score5 := 78

	sum := score1 + score2 + score3 + score4 + score5
	average := float64(sum) / 5
	fmt.Printf("总分：%d\n平均分：%.2f\n", sum, average)

	// 用数组实现相同的功能
	scores := [5]int{95, 91, 63, 86, 78}
	sum = 0
	for _, score := range scores {
		sum += score
	}
	average = float64(sum) / 5
	fmt.Printf("总分：%d\n平均分：%.2f\n", sum, average)
}
