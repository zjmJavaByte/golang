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

	//在格式化输出时，你可以使用 %t 来表示你要输出的值为布尔型
	fmt.Printf("值：%t \n", b)

	//注意事项
	//占用的存储空间是一个字节
	fmt.Printf("b占用的字节空间：%d \n", unsafe.Sizeof(b))
}
