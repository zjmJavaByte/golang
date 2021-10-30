package main

import (
	"container/list"
	"fmt"
	"unsafe"
)

func main() {

	/**
	练习 9.1
	使用 container/list 包实现一个双向链表，将 101、102 和 103 放入其中并打印出来。
	*/
	l := list.New() //New创建一个链表

	l.PushBack(101) //PushBack将一个值为v的新元素插入链表的最后一个位置，返回生成的新元素。
	l.PushBack(102)
	l.PushBack(103)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("值：%v 前：%v 后： %v \n", e.Value, e.Prev(), e.Next())
	}

	/*
		练习 9.2
		通过使用 unsafe 包中的方法来测试你电脑上一个整型变量占用多少个字节。
	*/

	a := 10
	fmt.Printf("%v \n", unsafe.Sizeof(a))
}
