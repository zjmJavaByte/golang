package main

import "fmt"

/*
变量使用的注意事项
*/
func main() {

	//一、 "该区域" 的值可以在 "同一类型" 的范围内不断变化
	var i int = 10
	fmt.Println(i)

	i = 100
	//i = 100.2 会报错，编译不通过
	fmt.Println(i)

}
