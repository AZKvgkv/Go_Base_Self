package main

import "fmt"
import "awesomeProject/src/CRM/dbutils"

func main() {
	fmt.Printf("你好 出现这个代表着main已经执行！")
	dbutils.AddNums()
}
