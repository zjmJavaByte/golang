package main

import (
	"fmt"
)

type B struct {
	thing int
}

func (b *B) change() { b.thing = 1 }

//func (b B) change() { b.thing = 1 }

func (b B) write() string {
	return fmt.Sprint(b)
}

func main() {

	/*
		1、当chang接受者是指针的时候，不管是调用者是指针还是值，都会改变原来的值
		   当chang接受者是值的时候，不管是调用者是指针还是值，都不会改变原来的值
		2、指针方法和值方法都可以在指针或非指针上被调用
	*/
	var b1 B // b1是值
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) // b2是指针
	b2.change()
	fmt.Println(b2.write())

}

/* 输出：
{1}
{1}
*/
