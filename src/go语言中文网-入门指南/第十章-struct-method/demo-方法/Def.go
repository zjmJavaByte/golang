package main

import (
	"fmt"
)

type Person struct {
	a, b int
}

type Student struct {
	age int
}

type intArr []int

type employee struct {
	salary float64
}

func main() {

	/*
			定义：
				Go 方法是作用在接收者（receiver）上的一个函数，接收者是某种类型的变量。因此方法是一种特殊类型的函数。
			接收者：
				1、任何类型都可以有方法，甚至可以是函数类型，可以是 int、bool、string 或数组的别名类型。
				但是接收者不能是一个接口类型，因为接口是一个抽象定义，但是方法却是具体实现
				2、接收者不能是一个指针类型，但是它可以是任何其他允许类型的指针
			Java与go中的方法区别：
				一个类型加上它的方法等价于面向对象中的一个类。一个重要的区别是：在 Go 中，类型的代码和绑定在它上面的方法的代码可以不放置在一起，它们可以存在在不同的源文件，唯一的要求是：它们必须是同一个包的。
		--------------------------------
					package main
					import "container/list"

					func (p *list.List) Iter() {
						// ...
					}

					func main() {
						lst := new(list.List)
						for _= range lst.Iter() {
						}
					}

		--------------------------------
				类型和作用在它上面定义的方法必须在同一个包里定义，这就是为什么不能在 int、float 或类似这些的类型上定义方法
					func (i int) name()  {

					}

					func (t time.Time) first3Chars() string {
						return time.LocalTime().String()[0:3]
					}

		---------------------------------
				将它作为匿名类型嵌入在一个新的结构体中
					type myTime struct {
		    			time.Time //anonymous field
					}

					func (t myTime) first3Chars() string {
						return t.Time.String()[0:3]
					}

	*/

	p := new(Person)
	p.a = 1
	p.b = 2
	add1 := p.add()
	add2 := p.addTwo(3)
	fmt.Println(add1)
	fmt.Println(add2)

	p2 := Person{1, 3}
	p2.a = 1
	p2.b = 2
	add3 := p2.add()
	add4 := p2.addTwo(5)
	fmt.Println(add3)
	fmt.Println(add4)

	arr := intArr{1, 2, 3}
	fmt.Println(arr.sum())

	ints := make(intArr, 0)
	ints = append(ints, 1)
	ints = append(ints, 1)
	ints = append(ints, 1)
	fmt.Println(ints.sum())

	e := new(employee)
	e.salary = 2.99
	raise := e.giveRaise(10)
	fmt.Println(raise)
}

/*
定义的格式：
func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
*/

/*
同一个接受者中，方法不能够重载
func (receiver Person) add() {

}

func (receiver Person) add(a int) {

}

不同的接受者，方法能够重载
func (receiver Person) add() {

}

func (receiver Student) add() {

}*/

func (this Person) add() int {
	return this.a + this.b
}

func (this Person) addTwo(c int) int {
	return this.a + this.b + c
}

func (this intArr) sum() (a int) {
	for _, v := range this {
		a += v
	}
	return
}

func (receiver employee) giveRaise(a float64) float64 {
	return receiver.salary * a
}
