package main

import "fmt"
import "time"

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c //此处在15秒之前没有链接通道，所以后面的主程序中个通道赋值是阻塞的
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
}
