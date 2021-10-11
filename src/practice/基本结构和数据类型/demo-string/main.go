package main

import (
	"fmt"
)

func main() {

	a := "1234\n56"
	b := "1234\r56"
	c := "123\t456"
	d := "`This is a raw string \n` 中的 `\n` 会被原样输出。"
	e := "123\\456"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	/**
	字符串拼接
	*/
	str := "Beginning of the string " +
		"second part of the string"
	fmt.Println(str)

	s := "hel" + "lo,"
	s += "world!"
	fmt.Println(s) //输出 “hello, world!”

	// TODO
	//使用函数拼接字符串
	//strings.Join()

	//使用字节缓冲（bytes.Buffer）拼接

}
