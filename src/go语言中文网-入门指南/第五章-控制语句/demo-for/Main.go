package main

import "fmt"

const (
	FIZZ     = 3
	BUZZ     = 5
	FIZZBUZZ = 15
)

func main() {

	/**
	基于计数器的迭代
	*/
	//for 初始化语句; 条件语句; 修饰语句 {}

	for i := 0; i < 10; i++ {

	}

	for i, j := 0, 10; i < j; i, j = i+1, j-1 {

		fmt.Println(i)
		fmt.Println(j)

	}

	// 使用 for 结构创建一个简单的循环。要求循环 15 次然后使用 fmt 包来打印计数器的值。
	for i := 0; i < 15; i++ {
		fmt.Printf("The counter is at %d\n", i)
	}
	// 2:使用 goto 语句重写循环，要求不能使用 for 关键字。
	i := 0
START:
	fmt.Printf("The counter is at %d\n", i)
	i++
	if i < 15 {
		goto START
	}

	/**
	G
	GG
	GGG
	GGGG
	GGGGG
	GGGGGG
	*/
	//使用 2 层嵌套 for 循环。
	for i := 0; i < 6; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("G")
		}
		fmt.Println()
	}
	g := "G"
	//使用一层 for 循环以及字符串截断。
	for i := 0; i < 6; i++ {
		fmt.Println(g)
		g += "G"
	}

	for i := 1; i <= 100; i++ {
		switch {
		case i%FIZZBUZZ == 0:
			fmt.Println("FizzBuzz")
		case i%FIZZ == 0:
			fmt.Println("Fizz")
		case i%BUZZ == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}

	/**
	  基于条件判断的迭代
	*/

	var j = 5

	for j >= 0 {
		j = j - 1
		fmt.Printf("The variable j is now: %d\n", i)
	}

	/**
	无限循环
	*/
	/*for {
	}*/

	/**
	 for-range 结构
	这是 Go 特有的一种的迭代结构，您会发现它在许多情况下都非常有用。它可以迭代任何一个集合（包括数组和 map，详见第 7 和 8 章）。
	语法上很类似其它语言中 foreach 语句，但您依旧可以获得每次迭代所对应的索引。一般形式为：for ix, val := range coll { }。

	要注意的是，val 始终为集合中对应索引的值拷贝，因此它一般只具有只读性质，对它所做的任何修改都不会影响到集合中原有的值
	（译者注：如果 val 为指针，则会产生指针的拷贝，依旧可以修改集合中的原值）。一个字符串是 Unicode 编码的字符（或称之为 rune）集合，因此您也可以用它迭代字符串：
	*/
	str := "123456"
	for pos, char := range str {
		fmt.Printf("索引 %d ASCII %d \n", pos, +char)
	}

	//练习 5.9 以下程序的输出结果是什么？
	for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}

	//无限自增
	/*for i := 0; ; i++ {
		fmt.Println("Value of i is now:", i)
	}*/

	//一直为0
	/*for i := 0; i < 3; {
		fmt.Println("Value of i:", i)
	}*/

	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
