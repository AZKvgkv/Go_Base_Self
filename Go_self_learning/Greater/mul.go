package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg sync.WaitGroup

// func hello(i int) {
// 	fmt.Printf("hello golang %v\n", i)
// 	defer wg.Done()// goroutine 结束计数器-1
// }

// func main() {
// 	for i := 0; i < 10; i++ {
// 		go hello(i)
//         wg.Add(1)// 启动一个 goroutine 计数器+1
//     }
// 	wg.Wait()// 等待所有 goroutine 结束
// 	fmt.Println("main function")
// }
