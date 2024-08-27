package main

import (
	"errors"
	"fmt"
)

func main() {
	err := show()
	if err != nil {
		fmt.Println("自定义错误：", err)
		panic(err)
	}
}

func show() (err error) {
	num1 := 10
	var num2 int
	fmt.Print("请输入除数：")
	_, err = fmt.Scanf("%d", &num2)
	if err != nil {
		return err
	}
	if num2 == 0 {
		return errors.New("除数不能为0")
	} else {
		result := num1 / num2
		fmt.Println(result)
	}
	return nil
}
