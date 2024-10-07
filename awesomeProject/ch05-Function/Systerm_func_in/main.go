package main

import "fmt"

func main() {
	/*
		1 len()函数
		2 new()函数:--分配内存，第一个参数为类型，返回值为指向该类型的新分配的零值的指针
		3 make()函数:--分配内存，第一个参数为类型，第二个参数为长度
	*/
	// 定义一个字符串
	str := "hello golang"
	fmt.Println(len(str))

	// 以int类型为例
	num := new(int)
	fmt.Printf("num_value:%p,&num_value:%d,num's type:--> %T,num_loc:%v \n ", num, *num, num, &num)

	// 以管道为例
	ch := make(chan int, 10)
	fmt.Printf("ch_value:%p,&ch_value:%d,ch's type:--> %T,ch_loc:%v", ch, &ch, ch, &ch)
	fmt.Println(ch)
}
