package main

import "fmt"

/*
转义字符

\t  一个制表位，实现对齐功能

\n  换行符

\\	一个\

\"	一个"

\r	一个回车

123     3453
-----------------------
123
3453
-----------------------
123\3453
-----------------------
123"3453
-----------------------
34535678

*/
func main() {

	fmt.Println("123\t3453")
	fmt.Println("-----------------------")
	fmt.Println("123\n3453")
	fmt.Println("-----------------------")
	fmt.Println("123\\3453")
	fmt.Println("-----------------------")
	fmt.Println("123\"3453")
	fmt.Println("-----------------------")
	//回车，从当前的最前面开始输出，覆盖掉以前的内容
	fmt.Println("12345678\r3453")

	fmt.Println("-----------------------")

	fmt.Println("123", "3453")
	//相当于
	fmt.Println("123" +
		"3453")

}
