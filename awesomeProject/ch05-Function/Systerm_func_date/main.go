package main

import (
	"fmt"
	"time"
)

func main() {
	// 时间和日期的函数，需要导入包
	now := time.Now()
	// Now() 返回值是一个结构体
	fmt.Println(now) //2023-12-25 21:05:54.7567601 +0800 CST m=+0.001754601

	// 将日期以年月日时分秒按照格式输出为字符串
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	// Sprintf 可以得到这个字符串，以便后续使用
	dateStr := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(dateStr)

	// 可以指定格式,其中每个数字都是固定的
	dateStr2 := now.Format("2006-01-02 15:04:05\n")
	fmt.Println(dateStr2)
}
