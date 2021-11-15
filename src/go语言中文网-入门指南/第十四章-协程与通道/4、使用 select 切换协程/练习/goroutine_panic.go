package main

import "fmt"

func main() {

	ch := make(chan int)

	go tel(ch)
	for i := range ch {
		fmt.Println(i)
	}
	//由于通道未关闭，导致主线程无法结束  fatal error: all goroutines are asleep - deadlock!
}

func tel(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}

}
