package main

import "fmt"

func main() {

	ch := make(chan int)

	go tel2(ch)
	for i := range ch {
		fmt.Println(i)
	}

}

func tel2(ch chan int) {
	for i := 0; i < 15; i++ {
		ch <- i
	}

	close(ch) //通道关闭后，主线程重的for循环会自动检测通道是否关闭，然后结束for循环

}
