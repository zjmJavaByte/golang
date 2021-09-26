package main

import (
	"fmt"
	"strconv"
)

/*
string转为基本数据类型
*/
func main() {

	var str string = "true"
	var b bool
	//ParseBool(str string) (bool, error)
	//func ParseBool(str string) (value bool, err error)
	//返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。
	//因为我只想获取到value bool ，不想获取err 所以我是用_忽略掉错误返回值
	b, _ = strconv.ParseBool(str)
	fmt.Printf("b type is %T b=%v \n", b, b)

	/*
		ParseInt(s string, base int, bitSize int) (i int64, err error)
		base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；
		bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
		返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。
	*/
	var str2 string = "1234"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("n1 type is %T b=%v \n", n1, n1)

	/*
		ParseFloat(s string, bitSize int) (float64, error)
		如果s合乎语法规则，函数会返回最为接近s表示值的一个浮点数（使用IEEE754规范舍入）。
		bitSize指定了期望的接收类型，32是float32（返回值可以不改变精确值的赋值给float32），64是float64；
		返回值err是*NumErr类型的，语法有误的，err.Error=ErrSyntax；结果超出表示范围的，返回值f为±Inf，err.Error= ErrRange。
	*/
	var str3 string = "1234.56"
	var n2 float64
	n2, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("n2 type is %T b=%v \n", n2, n2)
}
