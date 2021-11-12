package main

import "fmt"

func main() {
	ch := make(chan int)
	go data3(50, ch)
	for i := range ch {
		fmt.Println(i)
	}

}

func data3(a int, ch chan int) {
	i, j := 1, 1
	for s := 0; s < a; s++ {
		ch <- i
		i, j = j, j+i
	}
	close(ch)
}
