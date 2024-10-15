package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 1; i <= 5; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	time.Sleep(time.Second * 1)
}
