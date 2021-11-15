package main

import (
	"fmt"
	"os"
)

func main() {

	ch := make(chan int)
	done := make(chan bool)
	go tel3(ch, done)
	for true {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-done:
			os.Exit(0)
		}
	}

}

func tel3(ch chan int, done chan bool) {
	for i := 0; i < 15; i++ {
		ch <- i
	}

	done <- true

}
