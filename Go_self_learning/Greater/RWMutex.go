package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var (
// 	x=0
// 	wg sync.WaitGroup
// 	rw_lock sync.RWMutex
// )

// func read() {
// 	defer wg.Done()
// 	rw_lock.RLock()
// 	fmt.Println("Reading",x)
// 	time.Sleep(time.Millisecond)
// 	rw_lock.RUnlock()
// }
// func write(){
// 	defer wg.Done()
//     rw_lock.Lock()
// 	x+=1
//     fmt.Println("Writing",x)
//     time.Sleep(time.Millisecond*5)
//     rw_lock.Unlock()
// }
// func main() {
// 	start := time.Now()
// 	for i := 0; i < 10; i++ {
// 		go write()
// 		wg.Add(1)
// 	}
// 	time.Sleep(time.Second)
// 	for i := 0; i < 1000; i++ {
//         go read()
//         wg.Add(1)
//     }
// 	wg.Wait()
// 	fmt.Println("Time taken", time.Since(start))
// }