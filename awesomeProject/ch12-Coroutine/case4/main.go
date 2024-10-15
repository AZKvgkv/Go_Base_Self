package main

import (
	"fmt"
	"sync"
)

var totalNum int
var wg sync.WaitGroup
var look sync.Mutex

func add() {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		look.Lock()
		totalNum++
		look.Unlock()
	}
}
func sub() {
	defer wg.Done()
	for i := 0; i < 1000000; i++ {
		look.Lock()
		totalNum--
		look.Unlock()
	}
}
func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(totalNum)

}
