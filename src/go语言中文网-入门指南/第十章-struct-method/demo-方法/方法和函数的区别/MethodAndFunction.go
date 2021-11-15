package main

import "fmt"

/*
函数和方法的区别
*/
type Teacher struct {
	age int
}

func main() {

	/*
		区别
			1、调用方式
				函数将变量作为参数：Function1(recv)
				方法在变量上被调用：recv.Method1()
			2、改变参数的值
				接收者是指针时，方法可以改变接收者的值
				函数参数为指针的时候，可以改变参数的值
		位置：方法没有和数据定义（结构体）混在一起：它们是正交的类型；表示（数据）和行为（方法）是独立的。
	*/

	s := new(Teacher)
	s.age = 1
	s.prin()
	fmt.Println(s)
}

func (Teacher) prin() {
	//s.age = 10
}

func add() {

}
