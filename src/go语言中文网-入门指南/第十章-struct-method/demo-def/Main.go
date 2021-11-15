package main

import "fmt"

/*
结构体定义
*/

type Person struct {
	name string
	age  int
}

/*

struct  &获取到的是&{zjm 10}
基本数据类型  0xc000102058

*/
func main() {

	//使用 new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针
	p := new(Person)
	p.age = 20
	p.name = "朱健民"
	fmt.Printf("p指针为：%v \n", p)
	fmt.Printf("p类型为：%T \n", p)
	fmt.Printf("p值为：%v \n", *p)

	//声明 var t T 也会给 t 分配内存，并零值化内存，但是这个时候 t 是类型 T
	var p1 Person
	fmt.Printf("p1值为：%v \n", p1)
	fmt.Printf("p1类型为：%T \n", p1)
	//fmt.Printf("p1指针为：%v \n",*p1)  Invalid indirect of 'p1' (type 'Person')

	//初始化一个结构体实列
	ms := Person{"zjm", 10}
	fmt.Println(ms)
	fmt.Printf("ms类型为：%T \n", ms)
	add := &Person{"zjm", 10}
	// 此时add的类型是 *Person
	add.age = 20 //可以通过指针打点直接赋值或者取值
	fmt.Println(add)
	fmt.Printf("add类型为：%T \n", add) //*main.Person

	var sub Person
	sub = Person{"zjm", 10}
	sub.age = 30 //可以通过变量打点直接赋值或者取值
	fmt.Println(sub)
	fmt.Printf("sub类型为：%T \n", sub) //main.Person

	abb := 1
	fmt.Println(&abb)
}
