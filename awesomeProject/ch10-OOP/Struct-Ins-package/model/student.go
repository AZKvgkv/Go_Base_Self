package model

type Student struct {
	Name string
	Age  int
}

// 工厂模式
func NewStudent(name string, age int) *Student {
	return &Student{Name: name, Age: age}
}