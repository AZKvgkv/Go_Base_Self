package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open a file for writing
	file, err := os.OpenFile("ch11-File/write/output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// 写入文件操作：---> io 流 ---> 缓冲输出流
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("This is line " + fmt.Sprint(i+1) + "\n")
	}
	// 刷新缓冲区
	writer.Flush()
}
