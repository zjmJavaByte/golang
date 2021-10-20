package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s := "\u00ff\u754c"
	//获取一个字节数组的切片 c
	c := []byte(s)
	for i, b := range c {
		fmt.Printf("%d:%c \n", i, b)
	}
	fmt.Println()
	for i, c := range s {
		fmt.Printf("%d:%c \n", i, c)
	}

	s2 := "12345"
	var bytes = make([]byte, 0)
	copy(bytes, s2)
	fmt.Println(bytes)

	f := []int32(s)
	fmt.Printf("[]int32(s) %v \n", f)
	r := []rune(s)
	fmt.Printf("[]rune(s) %v \n", r)
	i := len([]int32(s))
	fmt.Printf("[]int32(s)长度是： %v \n", i)
	inString := utf8.RuneCountInString(s)
	fmt.Printf("utf8.RuneCountInString长度是： %v \n", inString)

	fmt.Println("--------------7.6.2 获取字符串的某一部分----------------------")
	str := "12345"
	substr := str[2:4]
	fmt.Println(substr)
}
