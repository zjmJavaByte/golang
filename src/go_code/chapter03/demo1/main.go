package main

import "fmt"

/*
变量的声明
*/

//定义全局变量一
var b1 = 100
var b2 = 100
var b3 = 100

//定义全局变量二
var (
	c1 = 200
	c2 = 200
	c3 = 200
)

func main() {

	//声明类型  var 变量名 类型 = 值
	var i int = 10
	fmt.Println(i)

	//类型推导
	var j = 10.1
	fmt.Println(j)

	//省略var 关键字
	name := "zjm"
	fmt.Println(name)

	//多变量声明一
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	//多变量声明二
	var s1, s2, s3 = 100, "zjm", 888
	fmt.Println(s1, s2, s3)

	//多变量声明三
	a1, a2, a3 := 100, "zjm", 888
	fmt.Println(a1, a2, a3)

	//定义全局变量一
	fmt.Println(b1, b2, b3)

	//定义全局变量二
	fmt.Println(c1, c2, c3)
}
