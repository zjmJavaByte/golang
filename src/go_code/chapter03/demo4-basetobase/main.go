package main

import "fmt"

/*
基本数据类型相互转换
*/
func main() {

	//int -->> float32
	var a int = 10
	var b float32 = float32(a)
	fmt.Printf("a=%v b=%v \n", a, b) //a=10 b=10

	//float64 -->> int
	var c float64 = 66.6
	var d int = int(c)
	fmt.Printf("c=%v d=%v \n", c, d) //c=66.6 d=66

	//------注意细节----------
	//虽然经过了转换，但是原先的类型还是不变的
	fmt.Printf("c type is %T \n", c) //c type is float64

	//大范围转位小范围，可能会溢出,只是转换的结果是按照溢出处理，和我们希望的结果不一样
	var e int64 = 66666
	var f int8 = int8(e) //-128～127
	fmt.Println(f)       //106

	//--------练习题-----------
	var n1 int32 = 12
	var n2 int64
	var n3 int8

	//Cannot use 'n1 + 20' (type int32) as the type int64
	//n2 = n1 + 20
	//n3 = n1 +20

	n2 = int64(n1) + 20
	n3 = int8(n1) + 20

	fmt.Println(n1, n2, n3)

}
