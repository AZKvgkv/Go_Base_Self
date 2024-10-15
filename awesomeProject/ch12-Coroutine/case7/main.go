package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	intChan := make(chan int, 10)
	wg.Add(2)
	go writeData(intChan)
	go readData(intChan)
	wg.Wait()

}

func writeData(intChan chan int) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		intChan <- i
		fmt.Printf("写入数据%d\n", i)
		time.Sleep(time.Second)
	}
	close(intChan)
}

func readData(intChan chan int) {
	defer wg.Done()
	for v := range intChan {
		fmt.Printf("读取数据%d\n", v)
		time.Sleep(time.Second)
	}
}
