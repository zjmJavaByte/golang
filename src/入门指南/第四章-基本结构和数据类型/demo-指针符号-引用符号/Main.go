package main

import "fmt"

func main() {
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ", a)                       //1
	fmt.Println("&a = ", &a)                     //0xc00000a098
	fmt.Println("*&a = ", *&a)                   //1
	fmt.Println("b = ", b)                       //0xc00000a098
	fmt.Println("&b = ", &b)                     //0xc000006028
	fmt.Println("*&b = ", *&b)                   //0xc00000a098
	fmt.Println("*b = ", *b)                     //1
	fmt.Println("c = ", c)                       //0xc000006028
	fmt.Println("*c = ", *c)                     //0xc00000a098
	fmt.Println("&c = ", &c)                     //0xc00000a098
	fmt.Println("*&c = ", *&c)                   //0xc000006028
	fmt.Println("**c = ", **c)                   //1
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c) //1
	fmt.Println("x = ", x)                       //1

}
