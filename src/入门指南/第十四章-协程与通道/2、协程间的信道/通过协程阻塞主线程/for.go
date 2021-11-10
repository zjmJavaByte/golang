package main

import "fmt"

func main() {

	ch := make([]chan int, 0)

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10000; i++ {
				fmt.Println(i)
				ch[i] <- i
			}
		}()
	}

	for i := range ch {
		<-ch[i]
	}

}
