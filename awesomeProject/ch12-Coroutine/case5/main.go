package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var lock sync.RWMutex

func read() {
	defer wg.Done()
	lock.RLock()

	fmt.Println("读")
	time.Sleep(time.Second)
	fmt.Println("读结束")
	lock.RUnlock()

}
func write() {
	defer wg.Done()
	lock.Lock()
	fmt.Println("写")
	time.Sleep(time.Second)
	fmt.Println("写结束")
	lock.Unlock()
}
func main() {
	wg.Add(6)
	for i := 1; i <= 5; i++ {
		go read()
	}
	go write()
	wg.Wait()
}
