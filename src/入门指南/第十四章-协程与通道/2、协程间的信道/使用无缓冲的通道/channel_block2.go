package main

import (
	"fmt"
	"time"
)

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
