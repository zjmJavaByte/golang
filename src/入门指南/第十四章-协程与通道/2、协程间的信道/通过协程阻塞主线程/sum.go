package main

import (
	"fmt"
)

func main() {

	ch := make(chan int)
	go sum(ch, 10000, 20000)
	for i := 0; i < 100000; i++ {
		fmt.Println(i)
	}
	<-ch

}

func sum(ch chan int, a, b int) {
	for i := 0; i < 10000; i++ {
		//fmt.Println(i)
	}
	ch <- a + b

}
