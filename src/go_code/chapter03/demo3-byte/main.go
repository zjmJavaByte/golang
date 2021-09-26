package main

import "fmt"

/*
字符类型
*/
func main() {

	//参考ASCII

	var a byte = '\t'
	b := '1'
	c := 'a'
	d := '{'
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//如果希望输出对应的字符，需要使用格式化输出
	fmt.Printf("a=%c b=%c c=%c d=%c  \n", a, b, c, d)

	//var e byte = '朱'   编译报错
	var e int = '朱'
	fmt.Printf("e=%c e对应的码值：%d \n", e, e) //e=朱 e对应的码值：26417

	//直接给某个变量赋值，然后格式化输出对应的unicode字符
	var f int = 122
	fmt.Printf("f=%c  \n", f) //z

}
