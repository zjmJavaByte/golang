package main

import "fmt"

/*

匿名字段
*/
type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b   int
	c   float32
	int // anonymous field
	//此处包含了一个结构体，其实相当于继承了 innerS ，所以outerS里面包含innerS的字段
	innerS //anonymous field
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10
	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	var o outerS
	o.b = 6
	o.c = 7.5
	o.int = 60
	o.in1 = 5
	o.in2 = 10
	o.innerS = innerS{5, 10}
	fmt.Printf("o.b is: %d\n", o.b)
	fmt.Printf("o.c is: %f\n", o.c)
	fmt.Printf("o.int is: %d\n", o.int)
	fmt.Printf("o.in1 is: %d\n", o.in1)
	fmt.Printf("o.in2 is: %d\n", o.in2)
	fmt.Printf("o.innerS is: %v\n", o.innerS)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)

	//于是可以得出一个结论：在一个结构体中对于每一种数据类型只能有一个匿名字段。
}
