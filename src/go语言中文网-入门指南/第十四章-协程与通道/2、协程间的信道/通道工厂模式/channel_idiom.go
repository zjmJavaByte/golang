package main

import (
	"fmt"
	"time"
)

/*
编程中常见的另外一种模式如下：不将通道作为参数传递给协程，而用函数来生成一个通道并返回（工厂角色）；函数内有个匿名函数被协程调用。

以下代码是此代码的优化
func main() {

	ch := make(chan int)
	go push2(ch)
	go sunk2(ch)
	time.Sleep(1e9)

}

func push2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func sunk2(ch chan int) {
	for true {
		fmt.Println(<-ch)
	}
}

*/
func main() {
	stream := pump()
	go suck(stream)
	time.Sleep(1e9)
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
