package main

import (
	"io"
	"os"
)

func main() {
	// 定义源文件
	file1Path := "ch11-File/write/output.txt"
	// 定义目标文件
	file2Path := "ch11-File/copy/demo.txt"

	// 打开源文件
	sourceFile, err := os.Open(file1Path)
	if err != nil {
		panic(err) // 处理打开文件的错误
	}
	defer sourceFile.Close() // 确保在函数结束时关闭文件

	// 创建目标文件
	destinationFile, err := os.Create(file2Path)
	if err != nil {
		panic(err) // 处理创建文件的错误
	}
	defer destinationFile.Close() // 确保在函数结束时关闭文件

	// 复制源文件内容到目标文件
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		panic(err) // 处理复制过程中的错误
	}
}

/*
下面是方式二
0666: 所有用户可读写，所有用户组可读写，所有其他用户可读写。
0644: 所有用户可读，所有用户组可读，所有其他用户可读。
0600: 所有用户可读，所有用户组不可读，所有其他用户不可读。
新版Go语言中 引入了 os.ModePerm ,所有平台都表示可读、可写、可执行权限（即0777）。
*/
// package main

// import "os"

// func main() {
// 	// 定义源文件
// 	file1Path := "ch11-File/write/output.txt"
// 	// 定义目标文件
// 	file2Path := "ch11-File/copy/demo1.txt"
// 	content, err := os.ReadFile(file1Path)
// 	if err != nil {
// 		panic(err) // 处理打开文件的错误
// 	}
// 	err = os.WriteFile(file2Path, content, 0644)
// 	if err != nil {
// 		panic(err) // 处理创建文件的错误
// 	}
// }
