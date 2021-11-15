package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go push(ch)
	fmt.Println(<-ch)
	time.Sleep(1e9)

}

func push(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
