package main

import (
	"fmt"
	"unsafe"
)

/*
整数数据类型
int8
*/
func main() {

	//有符号整型
	var i int //大小和系统有关
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64

	fmt.Println(i, i8, i16, i32, i64)

	//无符号整型
	var a uint //大小和系统有关
	var a8 uint8
	var a16 uint16
	var a32 uint32
	var a64 uint64
	fmt.Println(a, a8, a16, a32, a64)

	//与int32一样，表示一个Unicode码
	var b rune
	fmt.Println(b)

	//无符号 等价于uint8，当要存储字符时，选用byte
	var c byte = '1'
	fmt.Println(c)

	//--------整型的使用细节----------
	//n1是什么类型
	var n1 = 100
	//可以用Printf 做格式化输出
	fmt.Printf("n1 的类型是：%T \n", n1)

	//n2占用的字节大小和数据乐行
	var n2 int64 = 10
	fmt.Printf("n2 的类型是：%T  n2占用的字节数是 %d \n", n2, unsafe.Sizeof(n2))

	//保证程序运行正确的情况下，尽量使用占用空间小的数据
	var age byte = 90
	fmt.Println(age)
}
