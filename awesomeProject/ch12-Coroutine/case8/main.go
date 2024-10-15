package main

import "fmt"

func main() {

	// 声明一个只写管道
	intChan1 := make(chan<- int, 10)
	intChan1 <- 10
	fmt.Println(intChan1)

	// 声明一个只读管道
	intChan2 := make(<-chan int, 10)
	fmt.Println(intChan2)

}
