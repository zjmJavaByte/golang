package main

import "fmt"

func main() {

	var a = true

	if true {
		a = false
		fmt.Println(a)
	}
	fmt.Println(a)

}
