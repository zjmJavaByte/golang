package main

import (
	"bytes"
	"fmt"
	"math/rand"
)

func main() {

	/**
	定义格式：var slice1 []type = arr1[start:end]
	*/

	var arr1 [6]int
	//终止索引标识的项不包括在切片内   [2,5)
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	//切片（slice）是对数组一个连续片段的引用  改变数组中的值，切片中的值也随之更改了
	arr1[4] = 10000
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	//和数组不同的是，切片的长度可以在运行时修改，最小为 0 最大为相关数组的长度：切片是一个 长度可变的数组
	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	//容量:等于切片从第一个元素开始，到相关数组末尾的元素个数
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice beyond capacity
	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	//println(b[1:4]) //0,l,a
	//println(b[:2])  //g,o
	//println(b[2:])  //l,a,n,g
	//println(b[:])   //g,o,l,a,n,g
	bytess := b[2:]
	for i := 0; i < len(bytess); i++ {
		fmt.Printf("bytess at %d is %c\n", i, bytess[i])
	}

	fmt.Printf("-------------7.2.2 将切片传递给函数---------------")
	var arr = [5]int{0, 1, 2, 3, 4}
	sum(arr[:])

	fmt.Println("---------------7.2.3 用 make () 创建一个切片---------------")
	//定义
	//这里 len 是数组的长度并且也是 slice 的初始长度
	//var slice2 []type = make([]type, len)

	//slice1 := make([]type, len)

	slice2 := make([]int, 10)
	for i := 0; i < len(slice2); i++ {
		slice2[i] = i
	}
	//var slice3 []type = make([]type, len, cap)
	slice3 := make([]int, 50, 100)
	slice4 := new([100]int)[0:50]
	for i := 0; i < len(slice3); i++ {
		slice3[i] = i
	}
	for i := 0; i < len(slice4); i++ {
		slice4[i] = i
	}

	/**
	练习 7.4： fobinacci_funcarray.go: 为练习 7.3 写一个新的版本，主函数调用一个使用序列个数作为参数的函数，该函数返回一个大小为序列个数的 Fibonacci 切片。
	*/

	ints := fbnq(15)
	for ix, fib := range ints {
		fmt.Printf("The %d-th Fibonacci number is: %d\n", ix, fib)
	}

	fmt.Println("--------------7.2.4 new () 和 make () 的区别----------------")
	//new (T) 为每个新的类型 T 分配一片内存，初始化为 0 并且返回类型为 * T 的内存地址：这种方法 返回一个指向类型为 T，值为 0 的地址的指针
	//它适用于值类型如数组和结构体；它相当于 &T{}
	var p *[]int = new([]int) // *p == nil; with len and cap 0
	//p := new([]int)
	fmt.Printf("p-->>值是否为空%v len %v cap %v  类型 %T \n", *p == nil, len(*p), cap(*p), p)

	//make(T) 返回一个类型为 T 的初始值,它只适用于 3 种内建的引用类型：切片、map 和 channel
	var p2 = make([]int, 0)
	fmt.Printf("p2-->>值是否为空%v len %v cap %v 类型 %T \n", p2 == nil, len(p2), cap(p2), p2)

	/**
	  问题 7.3 给定 s := make([]byte, 5)，len (s) 和 cap (s) 分别是多少？s = s[2:4]，len (s) 和 cap (s) 又分别是多少？
	*/
	s := make([]byte, 5)
	fmt.Printf("s-->>len %v cap %v  \n", len(s), cap(s)) // 5  5
	s = s[2:4]
	fmt.Printf("s-->>len %v cap %v  \n", len(s), cap(s)) // 2 3

	/**
	问题 7.4 假设 s1 := []byte{'p', 'o', 'e', 'm'} 且 s2 := s1[2:]，s2 的值是多少？如果我们执行 s2[1] = 't'，s1 和 s2 现在的值又分别是多少？
	*/
	s1 := []byte{'p', 'o', 'e', 'm'}
	s2 := s1[2:] // e m
	fmt.Printf("%c \n", s2)
	s2[1] = 't' // 'p', 'o', 'e', 't'     'e', 't'
	fmt.Printf("%c \n", s2)
	fmt.Printf("%c \n", s1)

	fmt.Println("------------7.2.6 bytes 包---------------")
	/**
	定义：
	var buffer bytes.Buffer
	或者使用 new 获得一个指针：
	var r *bytes.Buffer = new(bytes.Buffer)
	或者通过函数：
	func NewBuffer(buf []byte) *Buffer
	*/
	var buffer bytes.Buffer
	for {
		if s, ok := getNextString(); ok { //method getNextString() not shown here
			buffer.WriteString(s)
		} else {
			break
		}
	}
	fmt.Print(buffer.String(), "\n")

	/**
	练习 7.5 给定切片 sl，将一个 []byte 数组追加到 sl 后面。写一个函数 Append(slice, data []byte) []byte，该函数在 sl 不能存储更多数据的时候自动扩容。
	*/
	var sl = make([]byte, 0)
	i := []byte{'a', 'b', 'a', 'b', 'a', 'b', 'a', 'a', 'a'}
	i2 := Append(sl, i)
	fmt.Printf("%c \n", i2)
	/**
	练习 7.6 把一个缓存 buf 分片成两个 切片：第一个是前 n 个 bytes，后一个是剩余的，用一行代码实现。
	*/
	n := 50
	buf := make([]int, 100)
	bytes1, bytes2 := buf[:n], buf[n:]
	fmt.Printf("%v \n", bytes1)
	fmt.Printf("%v \n", bytes2)
}

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

func fbnq(term int) (ints []int) {
	ints = make([]int, term)
	ints[0] = 1
	ints[1] = 1

	for i := 2; i < len(ints); i++ {
		ints[i] = ints[i-1] + ints[i-2]
	}
	return
}

func getNextString() (s string, bool bool) {
	str := "hello world"
	intn := rand.Intn(10)
	bool = true
	if intn == 5 {
		bool = false
	}
	s = str[intn-1 : intn]
	return
}

func Append(slice, data []byte) []byte {
	for _, datum := range data {
		slice = append(slice, datum)
		fmt.Println(cap(slice))
	}
	return slice
}
