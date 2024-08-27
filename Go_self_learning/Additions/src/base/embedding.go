package base

import (
	"fmt"
)

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base %d", b.num)
}

type container struct {
	base
	name string
}

func M_embedding() {
	co := container{
		base: base{num: 10},
		name: "container",
	}
	fmt.Printf("co={num: %v, name: %v}\n", co.num, co.name)
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	type describe interface {
		describe() string
	}

	var d describe = co
	fmt.Println("describe:", d.describe())
}
