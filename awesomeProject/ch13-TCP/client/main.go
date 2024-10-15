package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("客户端启动中...")
	coon, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("客户端启动失败，错误：", err)
		return
	}
	fmt.Println("客户端启动成功, 连接服务器成功, coon: ", coon)
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("客户端读取数据失败，错误：", err)
		return
	}
	n, err := coon.Write([]byte(str))
	if err != nil {
		fmt.Println("客户端发送数据失败，错误：", err)
		return
	}
	fmt.Println("客户端发送数据成功，发送了", n, "个字节")

}
