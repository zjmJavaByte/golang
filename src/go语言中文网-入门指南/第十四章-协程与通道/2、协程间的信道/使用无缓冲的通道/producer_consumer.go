package main

import "fmt"

/*
用这种习惯用法写一个程序，有两个协程，第一个提供数字 0，10，20，…90 并将他们放入通道，第二个协程从通道中读取并打印。main() 等待两个协程完成后再结束
*/
func main() {

	ch := make(chan int, 10)
	done := make(chan bool)
	go push5(ch)
	go get2(ch, done)
	<-done
}

func push5(ch chan int) {
	for i := 0; i <= 9; i++ {
		ch <- i * 10
	}
	close(ch)
}

func get2(ch chan int, done chan bool) {
	//for 循环的 range 语句可以用在通道 ch 上，便可以从通道中获取值
	for i := range ch {
		fmt.Println(i)
	}
	/*
		它从指定通道中读取数据直到通道关闭，才继续执行下边的代码。很明显，另外一个协程必须写入 ch（不然代码就阻塞在 for 循环了），而且必须在写入完成后才关闭
	*/

	done <- true
}
