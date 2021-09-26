package main

import (
	"fmt"
	"strconv"
)

/*
基本数据类型转为string
*/
func main() {

	/*-----------方式以------------*/
	var num1 int = 99
	var num2 float64 = 23.456
	var b bool = true
	var myChar byte = 'h'

	var str string

	/*
		str type string str="99"
		str type string str="23.456000"
		str type string str="true"
		str type string str="h"
	*/

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T str=%q \n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T str=%q \n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T str=%q \n", str, str)

	str = fmt.Sprintf("%c", myChar)
	fmt.Printf("str type %T str=%q \n", str, str)

	/*--------方式二----------*/

	var num3 int = 99
	var num4 float64 = 23.456
	var num5 int = 999
	var b2 bool = true

	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T str=%q \n", str, str) //str type string str="99"

	/*
		bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。

		fmt表示格式：'f'（-ddd.dddd）、'b'（-ddddp±ddd，指数为二进制）、
		'e'（-d.dddde±dd，十进制指数）、'E'（-d.ddddE±dd，十进制指数）、
		'g'（指数很大时用'e'格式，否则'f'格式）、'G'（指数很大时用'E'格式，否则'f'格式）。

		prec控制精度（排除指数部分）：对'f'、'e'、'E'，它表示小数点后的数字个数；
		对'g'、'G'，它控制总的数字个数。如果prec 为-1，则代表使用最少数量的、但又必需的数字来表示f。

		FormatFloat(f float64, fmt byte, prec, bitSize int)
	*/
	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T str=%q \n", str, str) //str type string str="23.4560000000"

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T str=%q \n", str, str) //str type string str="true"

	str = strconv.Itoa(num5)
	fmt.Printf("str type %T str=%q \n", str, str) //str type string str="999"
}
