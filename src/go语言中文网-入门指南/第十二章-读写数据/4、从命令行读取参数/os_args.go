// os_args.go
package main

import (
	"fmt"
	"os"
	"strings"
)

//D:\goProject\sggProject\src\入门指南\第十二章-读写数据\4、从命令行读取参数>go run os_args.go 朱健民 杨爱平
func main() {
	who := "Alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}
