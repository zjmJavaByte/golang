package main

import (
	"fmt"
	"time"
)

/*
使用通道让 main 程序等待协程完成，就是所谓的信号量模式
*/

func main() {

	ch := make(chan int)
	go sum(ch, 10000, 20000)
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	fmt.Println()
	fmt.Println(<-ch)

}

func sum(ch chan int, a, b int) {
	time.Sleep(1e9)
	ch <- a + b

}
