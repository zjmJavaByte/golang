package main

import (
	"fmt"
	"strings"
)

/**
将函数作为参数
*/
func main() {
	asciiOnly := func(c rune) rune {
		if c > 127 {
			return ' '
		}
		return c
	}
	fmt.Println(strings.Map(asciiOnly, "Jérôme Österreich"))
}
