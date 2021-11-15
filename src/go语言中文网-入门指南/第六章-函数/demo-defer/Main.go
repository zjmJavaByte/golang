package main

import (
	"fmt"
	"io"
	"log"
)

/**
defer
*/
func main() {
	function1()

	a()

	log.Println(123)

}

func function1() {
	fmt.Printf("In function1 at the top\n")
	defer function2() //关键字 defer 允许我们推迟到函数返回之前（或任意位置执行 return 语句之后）一刻才执行某个语句或函数
	fmt.Printf("In function1 at the bottom!\n")
}

func function2() {
	fmt.Printf("function2: Deferred until the end of the calling function!\n")
}

func a() {
	i := 0
	defer fmt.Printf("a %v \n", i)
	i++
	defer fmt.Printf("b %v \n", i)
	i++
	return
}

//当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）
func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}

//使用 defer 语句实现代码追踪
func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func c() {
	trace("a")
	defer untrace("a")
	fmt.Println("in a")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	c()
}

//使用 defer 语句来记录函数的参数与返回值
func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}
