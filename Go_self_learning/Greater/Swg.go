package main

// import (
//     "fmt"
//     "sync"
// )

// var wg sync.WaitGroup

// func hello(){
// 	fmt.Printf("Hello\n")
// 	defer wg.Done() // 把计数器-1
// }

// func main()  {
// 	wg.Add(1) // 把计数器+1
// 	go hello()
// 	fmt.Printf("World\n")
// 	wg.Wait() // 等待计数器归零
// }