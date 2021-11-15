package main

import "fmt"

type str string

var ss str = "朱健民"

func main() {

	/*
		匿名函数空接口定义
	*/
	f := func(any interface{}) {
		switch v := any.(type) {
		case bool:
			fmt.Printf("any %v is a bool type", v)
		case int:
			fmt.Printf("any %v is an int type", v)
		case float32:
			fmt.Printf("any %v is a float32 type", v)
		case string:
			fmt.Printf("any %v is a string type", v)
		case str:
			fmt.Printf("any %v is a special String!", v)
		default:
			fmt.Println("unknown type!")
		}
	}
	f(ss)
	fmt.Println()

	/*
		复制数据切片至空接口切片

		var dataSlice []myType = FuncReturnSlice()
			var interfaceSlice []interface{} = dataSlice

			可惜不能这么做，编译时会出错：cannot use dataSlice (type []myType) as type []interface { } in assignment。

			原因是它们俩在内存中的布局是不一样的（参考 官方说明）。

			必须使用 for-range 语句来一个一个显式地复制：

			var dataSlice []myType = FuncReturnSlice()
			var interfaceSlice []interface{} = make([]interface{}, len(dataSlice))
			for i, d := range dataSlice {
				interfaceSlice[i] = d
			}
	*/

	/*
		通用类型的节点数据结构
	*/
	a := NewNode(nil, nil, "left node")
	b := NewNode(nil, nil, "right node")
	root := NewNode(a, b, "root node")
	// make child (leaf) nodes:
	fmt.Printf("%v\n", root) // Output: &{0x125275f0 root node 0x125275e0}

	/*
		接口到接口
	*/

}

type Node struct {
	le   *Node
	data interface{}
	ri   *Node
}

func NewNode(le, ri *Node, v interface{}) *Node {
	return &Node{le, v, ri}
}
