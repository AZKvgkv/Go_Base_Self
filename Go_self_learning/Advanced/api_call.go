package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("nokia phone")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("iphone")
}
// func main() {
// 	var phone Phone
// 	phone = new(NokiaPhone)
// 	phone.call()
// 	phone = new(IPhone)
// 	phone.call()
// }