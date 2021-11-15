package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan string)

	go SendData(ch)
	go GetData(ch)

	time.Sleep(1e9) //单位：ns
}

func SendData(ch chan string) {
	/*
		这个操作符直观的标示了数据的传输：信息按照箭头的方向流动。
		流向通道（发送）
		ch <- int1 表示：用通道 ch 发送变量 int1（双目运算符，中缀 = 发送）
		从通道流出（接收），三种方式：
		int2 = <- ch 表示：变量 int2 从通道 ch（一元运算的前缀操作符，前缀 = 接收）接收数据（获取新值）；假设 int2 已经声明过了，如果没有的话可以写成：int2 := <- ch
	*/
	ch <- "1"
	ch <- "2"
	ch <- "3"
}

func GetData(ch chan string) {
	var input string
	//getData () 使用了无限循环：它随着 sendData () 的发送完成和 ch 变空也结束了。
	for true {
		input = <-ch
		fmt.Println(input)
	}
}
