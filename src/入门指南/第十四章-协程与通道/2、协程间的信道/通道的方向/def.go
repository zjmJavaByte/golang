package main

func main() {

	/*
		var send_only chan<- int         // channel can only receive data
		var recv_only <-chan int        // channel can only send data
	*/
	var c = make(chan int) // bidirectional
	go source(c)
	go sink(c)
}

func source(ch chan<- int) {
	for {
		ch <- 1
	}
}

func sink(ch <-chan int) {
	for {
		<-ch
	}
}
