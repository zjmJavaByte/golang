package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

}

/*func (st *Stack) Pop() int {
	v := 0
	for ix := len(st) - 1; ix >= 0; ix-- {
		if v = st[ix]; v != 0 {
			st[ix] = 0
			return v
		}
	}
	return v
}*/

/**
6.2.1 按值传递（call by value） 按引用传递（call by reference）
*/
//1、值传递
//Go 默认使用按值传递来传递参数，也就是传递参数的副本。函数接收参数副本之后，
//在使用变量的过程中可能对副本的值进行更改，但不会影响到原来的变量，
//比如 Function(arg1)

//2、引用传递
//如果你希望函数可以直接修改参数的值，而不是对参数的副本进行操作，你需要将参数的地址（变量名前面添加 & 符号，比如 &variable）
//传递给函数，这就是按引用传递，
//比如 Function(&arg1)，此时传递给函数的是一个指针

//注意事项：指针也是变量类型，有自己的地址和值，通常指针的值指向一个变量的地址。所以，按引用传递也是按值传递

//3、返回多个值
//如果一个函数需要返回四到五个值，我们可以传递一个切片给函数（如果返回值具有相同类型）或者是传递一个结构体（如果返回值具有不同的类型）。因为传递一个指针允许直接修改变量的值，消耗也更少。

//4、问题：一下两个函数有什么不同
/*func DoSomething(a *A) {  传递的是指针
	b = a
}*/

/*(B) func DoSomething(a A) {//传递的是赋值后的值
	b = &a     此处获取的是复制后的值得地址
}*/

/**
6.2.2 命名的返回值（named return variables）
*/

//练习题
/*
1、编写一个函数，接收两个整数，然后返回它们的和、积与差。编写两个版本，一个是非命名返回值，一个是命名返回值。
func calculate(a int,b int) (int,int,int)  {
	return a + b, a * b, a - b
}
*/
func calculate(a int, b int) (x int, y int, z int) {
	x = a + b
	y = a * b
	z = a - b
	return
}

/**
编写一个名字为 MySqrt 的函数，计算一个 float64 类型浮点数的平方根，如果参数是一个负数的话将返回一个错误。编写两个版本，一个是非命名返回值，一个是命名返回值。
*/

/*func MySqrt(a float64) (float64,error) {
	if a < 0 {
		return math.NaN(),errors.New("错误")
	}
	return math.Sqrt(a),nil
}*/
func MySqrt(a float64) (b float64, err error) {
	if a < 0 {
		b = math.NaN()
		err = errors.New("错误")
	} else {
		b = math.Sqrt(a)
	}
	return
}

/**
6.2.3 空白符（blank identifier）
*/

func getResult() {
	var i1 int
	var f1 float32
	i1, _, f1 = ThreeValues() //此处不接受值
	fmt.Printf("The int: %d, the float: %f \n", i1, f1)
}

func ThreeValues() (int, int, float32) {
	return 5, 6, 7.5
}

/**
6.2.4 改变外部变量（outside variable）
*/

// this function changes reply:
//传递的是指针，通过指针改变指针所对应的值
func Multiply(a, b int, reply *int) {
	*reply = a * b
}
