package main

import (
	"fmt"
	"time"
)

func main() {

	/*
		- 切片、map 和 channel，使用 make
		- 数组、结构体和所有的值类型，使用 new
	*/
	tick := time.Tick(1e8)
	boom := time.After(5e8)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}
}
