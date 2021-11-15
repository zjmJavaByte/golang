package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Println("Starting the server ...")

	//创建listen监听器   监听来自 localhost 即 IP 地址为 127.0.0.1 端口为 50000 基于 TCP 协议
	listen, err := net.Listen("tcp", "localhost:50000")

	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	}
	//坚挺来自客户端连接
	for i := 0; ; i++ {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		go doServerStuff(conn, i)
	}

}

func doServerStuff(conn net.Conn, i int) {
	//conn.RemoteAddr()
	fmt.Println("执行 doServerStuff 程序：", i)
	for true {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error read", err.Error())
			return // 终止程序
		}
		fmt.Printf("Received data: %v", string(buf[:len]))

	}

}
