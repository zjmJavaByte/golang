package main

import "fmt"

/**
递归
*/
func main() {
	result := 0
	n := 0
	for i := 0; i <= 10; i++ {
		result, n = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", n, result)
	}
	fmt.Println()
	print(1)

	fmt.Println()
	for i := uint64(0); i < uint64(30); i++ {
		fmt.Printf("Factorial of %d is %d\n", i, Factorial(i))
	}

}

//斐波那契数列
func fibonacci(n int) (res int, i int) {
	if n <= 1 {
		res = 1
	} else {
		res1, _ := fibonacci(n - 1)
		res2, _ := fibonacci(n - 2)
		res = res1 + res2
	}
	i = n
	return
}

//使用递归函数从 10 打印到 1
func print(a int) {
	if a > 10 {
		return
	}
	print(a + 1)
	fmt.Println(a)
}

//实现一个输出前 30 个整数的阶乘的程序。
/*func Factorial(n uint64) uint64 {
	if n > 0 {
		return n * Factorial(n-1)
	}
	return 1
}*/

func Factorial(n uint64) (a uint64) {
	a = 1
	if n > 0 {
		a = n * Factorial(n-1)
		return
	}
	return
}
