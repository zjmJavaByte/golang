package main

import (
	"fmt"
)

/*
使用main阻塞主线程
*/
func main() {

	//ch := make(chan int)
	go push()
	//time.Sleep(1e9)
	{
	}
}

func push() {
	for i := 0; ; i++ {
		fmt.Println(i)
	}
}
