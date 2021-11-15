package main

import "fmt"

/*
用这种习惯用法写一个程序，开启一个协程来计算 2 个整数的合并等待计算结果并打印出来
*/
func main() {
	ch := make(chan int)
	go push4(ch, 1, 2)
	fmt.Println(<-ch)

}

func push4(ch chan int, a, b int) {
	ch <- a + b

}
