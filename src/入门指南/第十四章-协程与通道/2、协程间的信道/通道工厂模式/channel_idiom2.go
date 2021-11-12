package main

import (
	"fmt"
	"time"
)

func main() {

	get(pump2())
	time.Sleep(1e9)
}

func pump2() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch

}

func get(ch chan int) {
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()
}
