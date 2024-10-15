package main

import "fmt"

type CInterface interface {
	MethodC()
}
type BInterface interface {
	MethodB()
}
type AInterface interface {
	CInterface
	BInterface
	MethodA()
}

type Stu struct {
}

func (s *Stu) MethodA() {
	println("MethodA")
}
func (s *Stu) MethodB() {
	println("MethodB")
}
func (s *Stu) MethodC() {
	println("MethodC")
}

type E interface {
}

func main() {
	var s Stu
	var a AInterface = &s
	a.MethodA()
	a.MethodB()
	a.MethodC()

	var e E = &s
	fmt.Println(e)
	var e2 interface{} = &s
	fmt.Println(e2)
	var num float64 = 3.14
	var e3 interface{} = num
	fmt.Println(e3)
}
