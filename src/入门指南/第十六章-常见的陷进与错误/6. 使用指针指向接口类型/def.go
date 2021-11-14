package main

func main() {

	/*

		查看如下程序：nexter 是一个接口类型，并且定义了一个 next() 方法读取下一字节。
		函数 nextFew1 将 nexter 接口作为参数并读取接下来的 num 个字节，并返回一个切片：这是正确做法。
		但是 nextFew2 使用一个指向 nexter 接口类型的指针作为参数传递给函数：当使用 next() 函数时，
		系统会给出一个编译错误：n.next undefined (type *nexter has no
		field or method next) （译者注：n.next 未定义（*nexter 类型没有 next 成员或 next 方法））

		例 16.1 pointer_interface.go (不能通过编译):

		package main
		import (
		    “fmt”
		)
		type nexter interface {
		    next() byte
		}
		func nextFew1(n nexter, num int) []byte {
		    var b []byte
		    for i:=0; i < num; i++ {
		        b[i] = n.next()
		    }
		    return b
		}
		func nextFew2(n *nexter, num int) []byte {
		    var b []byte
		    for i:=0; i < num; i++ {
		        b[i] = n.next() // 编译错误:n.next未定义（*nexter类型没有next成员或next方法）
		    }
		    return b
		}
		func main() {
		    fmt.Println("Hello World!")
		}
		永远不要使用一个指针指向一个接口类型，因为它已经是一个指针。

	*/

}
