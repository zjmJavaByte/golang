package main

import "fmt"

func main() {
	ch := make(chan int)
	go get(ch)
	for i := range ch {
		fmt.Println(i)
	}
}

func get(ch chan int) {
	for i := 1; i < 25; i++ {
		ch <- data2(i)
	}
	close(ch)
}

func data2(a int) (i int) {
	if a <= 1 {
		return a
	} else {
		i := data2(a-1) + data2(a-2)
		return i
	}
}
