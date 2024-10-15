package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 100)
	go func() {
		time.Sleep(time.Second)
		intChan <- 12
	}()
	stringChan := make(chan string, 100)
	go func() {
		time.Sleep(time.Second * 2)
		stringChan <- "hello"
	}()

	select {
	case v := <-intChan:
		fmt.Println(v)
	case v := <-stringChan:
		fmt.Println(v)
	default:
		fmt.Println("no data")
	}
}
