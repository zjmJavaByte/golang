package main

import (
	"fmt"
	"time"
)

/*
练习：解释为什么下边这个程序会导致 panic：所有的协程都休眠了 - 死锁！
主线程运行到 out <- 2 这里时，因为通道是无缓冲的且无接受者，所以被阻塞，此时程序就会报错
*/

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)
	//out := make(chan int, 1)  //solution 1
	out <- 2
	//go func(out)  solution 2
	go f1(out)

	time.Sleep(1e9)
}
