package main

import (
	"fmt"
	"strings"
)

func main() {

	func() {
		sum := 0
		for i := 1; i <= 1e6; i++ {
			sum += i
		}
	}()

	a := func() {
		sum := 0
		for i := 1; i <= 1e6; i++ {
			sum += i
		}
	}
	fmt.Printf(" - a is of type %T \n", a)

	fmt.Println(f())

	var g int
	go func(i int) {
		s := 0
		for j := 0; j < i; j++ {
			s += j
		}
		g = s
		fmt.Printf("%d \n", g)
	}(1000) // Passes argument 1000 to the function literal.

	fmt.Println()

	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	bmp := addBmp("file")   // returns: file.bmp
	jpeg := addJpeg("file") // returns: file.jpeg
	fmt.Println(bmp)
	fmt.Println(jpeg)
}

//函数 f 返回时，变量 ret 的值是什么
//变量 ret 的值为 2，因为 ret++ 是在执行 return 1 语句后发生的。
func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

/*func main() {
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
结果：1 - 21 - 321
三次调用函数 f 的过程中函数 Adder () 中变量 delta 的值分别为：1、20 和 300。
我们可以看到，在多次调用中，变量 x 的值是被保留的，即 0 + 1 = 1，然后 1 + 20 = 21，最后 21 + 300 = 321：
闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量
*/

/**
不使用递归但使用闭包改写第 6.6 节中的斐波那契数列程序

func main() {
	f := fib()
	// Function calls are evaluated left-to-right.
	// println(f(), f(), f(), f(), f())
	for i := 0; i <= 9; i++ {
		println(i+2, f())
	}
}

func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}
*/

//工厂函数:一个返回值为另一个函数的函数

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
