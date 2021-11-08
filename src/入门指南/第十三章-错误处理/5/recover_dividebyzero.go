package main

import "fmt"

func main() {

	fmt.Println("开始")
	test()
	fmt.Println("结束")
}

func test() {

	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("panic %v\r\n", err)
		}
	}()

	badCall()
	fmt.Printf("After bad call\r\n")
}

func badCall() {
	a, b := 10, 0
	n := a / b
	fmt.Println(n)
}
