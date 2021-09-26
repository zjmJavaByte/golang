package main

import (
	"fmt"
	"unsafe"
)

/*
浮点数据类型
浮点型 = 符号为 + 指数位 + 尾数位
*/
func main() {

	//单精度  4字节
	var price float32 = 33.6
	fmt.Println(price)

	//双精度  8字节
	var money float64 = 33.6
	fmt.Println(money)

	//浮点型默认声明位float64类型
	num1 := 64.1
	fmt.Printf("num1的类型位：%T 占用字节大小：%d \n", num1, unsafe.Sizeof(num1))

	//可以省略整数位的0
	num2 := .123
	fmt.Printf("num2的类型位：%T 占用字节大小：%d \n", num2, unsafe.Sizeof(num2))

	//科学计数法
	num3 := 5.1234e2  //相当于 5.1234 * 10 的二次方
	num4 := 5.1234e-2 //相当于 5.1234 / 10 的二次方
	fmt.Println(num3)
	fmt.Println(num4)

}
