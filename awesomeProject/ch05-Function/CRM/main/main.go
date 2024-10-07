package main

import (
	"awesomeProject/ch05-Function/CRM/dbutils"
	"fmt"
)

func main() {
	fmt.Printf("你好 出现这个代表着main已经执行！")
	dbutils.AddNums()
}