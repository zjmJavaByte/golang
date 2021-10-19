package main

import "fmt"

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice beyond capacity
	//slice1 = slice1[0:7 ] // panic: runtime error: slice bound out of range

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	//println(b[1:4]) //0,l,a
	//println(b[:2])  //g,o
	//println(b[2:])  //l,a,n,g
	//println(b[:])   //g,o,l,a,n,g
	bytes := b[2:]
	for i := 0; i < len(bytes); i++ {
		fmt.Printf("bytes at %d is %c\n", i, bytes[i])
	}

	fmt.Printf("7.2.2 将切片传递给函数")
	var arr = [5]int{0, 1, 2, 3, 4}
	sum(arr[:])

	fmt.Println("7.2.3 用 make () 创建一个切片")
	//定义
	//var slice1 []type = make([]type, len)
	//slice1 := make([]type, len)
}

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
