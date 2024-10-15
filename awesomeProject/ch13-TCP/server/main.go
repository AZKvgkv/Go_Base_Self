package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器端启动中...")
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("服务器端启动失败，错误：", err)
		return
	}
	fmt.Println("服务器端启动成功, 监听端口成功, listen: ", listen)

	// 监听成功之后，循环等待客户端的连接
	for {
		fmt.Println("等待客户端的连接...")

		coon, err := listen.Accept()
		if err != nil {
			fmt.Println("客户端的等待失败", err)
		} else {
			fmt.Printf("客户端的等待成功, coon:%v, 接收到的客户端信息：%v \n", coon, coon.RemoteAddr().String())
		}
		// 准备一个协程，专门处理客户端的请求
		go processRequest(coon)
	}
}
func processRequest(coon net.Conn) {
	defer coon.Close()
	for {
		// 创建一个切片，用于存储客户端发送的数据
		buf := make([]byte, 1024)

		// 读取客户端发送的数据
		n, err := coon.Read(buf)
		if err != nil {
			fmt.Println("客户端的读取失败，错误：", err)
			return
		}
		fmt.Println("客户端的读取成功，读取了", n, "个字节")
		fmt.Println("客户端的读取成功，读取了", string(buf[:n]))
	}
}
