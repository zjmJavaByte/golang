package main

import (
	"fmt"
	"time"
)

func main() {

	//函数通道
	funcChan := make(chan func())
	go push(funcChan)
	go get(funcChan)
	time.Sleep(1e9)
}

func push(ch chan func()) {
	ch <- func() {
		fmt.Println(123456)
	}

}

func get(ch chan func()) {
	v := <-ch
	v()
}
