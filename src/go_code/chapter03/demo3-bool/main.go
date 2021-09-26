package main

import (
	"fmt"
	"unsafe"
)

/*
布尔类型
*/
func main() {

	a := true
	fmt.Println(a)

	var b bool = false
	fmt.Println(b)

	//注意事项
	//占用的存储空间是一个字节
	fmt.Printf("b占用的字节空间：%d \n", unsafe.Sizeof(b))
}
