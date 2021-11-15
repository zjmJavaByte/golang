package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Go 语言的函数经常使用两个返回值来表示执行是否成功：返回某个值以及 true 表示成功；返回零值（或 nil）和 false 表示失败
	atoi, err := strconv.Atoi("1q")
	if err != nil {
		fmt.Println(err)
		//return
	}
	// 未发生错误，继续执行：
	fmt.Println(atoi)

}
