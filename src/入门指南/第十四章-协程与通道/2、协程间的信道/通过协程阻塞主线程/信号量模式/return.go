package main

import (
	"fmt"
	"time"
)

/*
有返回值的阻塞主线程
*/
func main() {

	ch := make(chan int)

	go doSomething(ch)
	fmt.Println("开始做某些事情")
	time.Sleep(1e9)
	fmt.Println("结束做某些事情")
	result := <-ch
	fmt.Println(result)
}

func doSomething(ch chan int) {
	time.Sleep(3e9)
	ch <- 1
}
