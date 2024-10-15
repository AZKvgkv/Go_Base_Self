package main

type AInterface interface {
	MethodA()
}

type BInterface interface {
	MethodB()
}

type Stu struct {
	AInterface
	BInterface
}

func (s *Stu) MethodA() {
	println("MethodA called")
}

func (s *Stu) MethodB() {
	println("MethodB called")
}
func main() {
	a := &Stu{}
	b := &Stu{}
	a.MethodA()
	b.MethodB()
}
