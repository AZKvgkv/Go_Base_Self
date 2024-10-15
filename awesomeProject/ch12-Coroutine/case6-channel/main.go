package main

import "fmt"

func main() {
	intChan := make(chan int, 10)
	fmt.Println(intChan)
	intChan <- 12
	intChan <- 73
	num := 41
	intChan <- num
	fmt.Printf("channel len = %d, cap = %d\n", len(intChan), cap(intChan))

	num1 := <-intChan
	fmt.Println(num1)
	fmt.Printf("channel len = %d, cap = %d\n", len(intChan), cap(intChan))

	close(intChan)
	num2 := <-intChan
	fmt.Println(num2)
}
