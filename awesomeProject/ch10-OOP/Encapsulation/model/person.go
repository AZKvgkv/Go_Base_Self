package model

type person struct {
	Name string
	age  int // 其他包不能访问到这个字段
}

func NewPerson(name string) *person {
	return &person{Name: name}
}
// 定义set和get方法，对age字段进行封装
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	}else{
		panic("年龄不合法")
	}
}

func (p *person) GetAge() int {
	return p.age
}