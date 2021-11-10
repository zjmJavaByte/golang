package main

import (
	"fmt"
	"time"
)

/*
练习：解释为什么下边这个程序会导致 panic：所有的协程都休眠了 - 死锁！
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
