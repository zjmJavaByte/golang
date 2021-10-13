package main

import "fmt"

/**
指针
*/
func main() {

	//Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
	var i1 = 6
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

	//指针 intP
	var intP *int
	//当一个指针被定义后没有分配到任何变量时，它的值为 nil
	fmt.Println(intP)
	//intP 存储了 i1 的内存地址；它指向了 i1 的位置，它引用了变量 i1
	intP = &i1
	fmt.Printf("it's location in memory: %p\n", intP)

	//符号 * 可以放在一个指针前，如 *intP，那么它将得到这个指针指向地址上所存储的值；这被称为反引用（或者内容或者间接引用）操作符
	fmt.Println(*intP)

	fmt.Println(*(&i1))
	fmt.Println(i1 == *(&i1))

	//通过对 *p 赋另一个值来更改 “对象”，这样 s 也会随之更改。
	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string

	//你不能得到一个文字或常量的地址
	const i = "1"
	//ptr := &i   Cannot take the address of 'i'
	//ptr := &1   Cannot take the address of '1'
	//ptr := &"1"   Cannot take the address of '"1"'

	//对一个空指针的反向引用是不合法的，并且会使程序崩溃
	//var a *int = nil
	//*a = 0
	//fmt.Println(a)
}
