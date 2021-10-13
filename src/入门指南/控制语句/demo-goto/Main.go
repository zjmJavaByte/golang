package main

import "fmt"

func main() {

	/**
	continue 语句指向 LABEL1，当执行到该语句的时候，就会跳转到 LABEL1 标签的位置。

	您可以看到当 j==4 和 j==5 的时候，没有任何输出：标签的作用对象为外部循环，因此 i 会直接变成下一个循环的值，而此时 j 的值就被重设为 0，即它的初始值。
	*/
LABEL1:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				continue LABEL1
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	/**
	如果将 continue 改为 break，则不会只退出内层循环，而是直接退出外层循环了
	*/
LABEL2:
	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j == 4 {
				break LABEL2
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	//还可以使用 goto 语句和标签配合使用来模拟循环
	i := 0
HERE:
	print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE

	/**
	特别注意 使用标签和 goto 语句是不被鼓励的：它们会很快导致非常糟糕的程序设计，而且总有更加可读的替代方案来实现相同的需求。

	一个建议使用 goto 语句的示例会在第 15.1 章的 simple_tcp_server.go 中出现：示例中在发生读取错误时，使用 goto 来跳出无限读取循环并关闭相应的客户端链接。

	定义但未使用标签会导致编译错误：label … defined and not used。

	如果您必须使用 goto，应当只使用正序的标签（标签位于 goto 语句之后），但注意标签和 goto 语句之间不能出现定义新变量的语句，否则会导致编译失败。

	*/

}
