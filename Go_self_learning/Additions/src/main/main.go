package main

import (
	"fmt"
	"src/base"
)
var value = "github"
func main()  {
	fmt.Printf("Hello, %v world!\n", value)
	base.M_closure()
	base.M_embedding()
	base.M_generic()
	base.M_goroutine()

	base.M_sort()
	base.M_func_sort()

	base.Text_template()
}