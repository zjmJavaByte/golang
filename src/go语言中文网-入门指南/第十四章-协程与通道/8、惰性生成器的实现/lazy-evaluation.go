package main

import "fmt"

var resume chan int

func main() {
	resume = integer()
	fmt.Println(<-resume)
	fmt.Println(<-resume)
	fmt.Println(<-resume)
	fmt.Println(<-resume)
	fmt.Println(<-resume)
}

func integer() chan int {
	ch := make(chan int)
	count := 0
	go func() {
		for true {
			ch <- count
			count++
		}
	}()
	return ch
}
