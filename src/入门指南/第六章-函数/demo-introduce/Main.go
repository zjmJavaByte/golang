package main

import "fmt"

func main() {

	person := Person{
		"123",
		12,
	}
	/**
	1、参数的传递与Java比较
	*/
	//1、括号里的是被调用函数的实参
	//2、函数被调用的时候，这些实参将被复制（简单而言）然后传递给被调用函数
	//3、Java是引用传递（复制）
	a := 1
	first(person, a)
	second(person, a)

}

type Person struct {
	name string
	age  int
}

//形参 :1
func first(a Person, b int) {
	a.name = "456"
	var pstr *Person
	pstr = &a
	var pstr2 *int
	pstr2 = &b
	fmt.Println(a)             //{456 12}
	fmt.Printf("%p \n", pstr)  //0xc000004078
	fmt.Printf("%p \n", pstr2) //0xc000128058
}

func second(a Person, b int) {
	var pstr *Person
	pstr = &a
	var pstr2 *int
	pstr2 = &b
	fmt.Println(a)             //{123 12}
	fmt.Printf("%p \n", pstr)  //0xc0000040a8
	fmt.Printf("%p \n", pstr2) //0xc000128058
}

/**
2、函数可以将其他函数调用作为它的参数，只要这个被调用函数的返回值个数、返回值类型和返回值的顺序与调用函数所需求的实参是一致的
假设 f1 需要 3 个参数 f1(a, b, c int)，同时 f2 返回 3 个参数 f2(a, b int) (int, int, int)，就可以这样调用 f1：f1(f2(a, b))
*/

/**
3、函数重载
指的是可以编写多个同名函数，只要它们拥有不同的形参 / 或者不同的返回值，在 Go 里面函数重载是不被允许的。这将导致一个编译错误
*/
//'second' redeclared in this package
/*func second(a Person,b Person) {
	var pstr = &a
	fmt.Println(a)//{123 12}
	fmt.Printf("%p \n",pstr)//0xc0000040a8
}*/
