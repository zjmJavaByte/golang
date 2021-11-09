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
