package main

import "fmt"

func main() {
	var C1 Classmate
	C1.Name = "吴双"
	C1.Age = 18
	C1.Sex = "女"

	var C2 = Classmate{}
	C2.Name = "郑国富"
	C2.Age = 20
	C2.Sex = "男"

	var C3 = new(Classmate)
	(*C3).Name = "陈诚"
	(*C3).Age = 19
	C3.Sex = "女" // 为了简化也可以这样赋值

	var C4 = &Classmate{}
	(*C4).Name = "周莹莹"
	(*C4).Age = 20
	C4.Sex = "女" // 同样为了简化也可以这样赋值

	var C5 = &Classmate{"徐龙", 22, "男"}

	fmt.Println(C1, C2, *C3, *C4, *C5)
}

type Classmate struct {
	Name string
	Age  int
	Sex  string
}
