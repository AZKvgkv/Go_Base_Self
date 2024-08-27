package main

import "fmt"

func main() {
	//获取终端输入，方式1：scanln;方式2：scanf
	var (
		age   int
		name  string
		score float32
		isVIP bool
	)
	//fmt.Println("你好")
	//fmt.Println("请输入学生的年龄：")
	//_, err := fmt.Scanln(&age)
	//if err != nil {
	//	return
	//}
	//fmt.Println("请输入学生的姓名：")
	//_, err = fmt.Scanln(&name)
	//if err != nil {
	//	return
	//}
	//fmt.Println("请输入学生的成绩：")
	//_, err = fmt.Scanln(&score)
	//if err != nil {
	//	return
	//}
	//fmt.Println("请输入学生的VIP状态：")
	//_, err = fmt.Scanln(&isVIP)
	//if err != nil {
	//	return
	//}
	//// fmt.Scanln(&age)也是可以的
	//fmt.Printf("学生的年龄是：%v,姓名为：%v成绩为：%v,是否为VIP：%v\n", age, name, score, isVIP)
	fmt.Println("请录入学生的年龄， 姓名， 成绩， 是否为VIP， 并使用空格进行分隔")
	_, _ = fmt.Scanf("%d %s %f %t", &age, &name, &score, &isVIP)
	fmt.Printf("学生的年龄是：%v,姓名为：%v成绩为：%v,是否为VIP：%v\n", age, name, score, isVIP)

}
