package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() { //主线程
	go test() // 开启一个协程
	for i := 1; i <= 3; i++ {
		fmt.Println("hello world + " + strconv.Itoa(i))
		// 延迟1秒
		time.Sleep(time.Second)
	}

}
func test() { //子线程
	for i := 1; i <= 3; i++ {
		fmt.Println("hello golang + " + strconv.Itoa(i))
		// 延迟1秒
		time.Sleep(time.Second)
	}
}
