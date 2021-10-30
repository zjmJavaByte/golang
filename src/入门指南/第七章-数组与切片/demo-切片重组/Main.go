package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	slice := arr[0:3]
	//slice = slice[0:4]
	//slice = slice[0:5]
	for i := range slice {
		fmt.Printf("%d \n", slice[i])
	}

	fmt.Println("------------------")

	slice1 := make([]int, 0, 10)
	// load the slice, cap(slice1) is 10:
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	/**
	问题 7.7
	1) 如果 s 是一个切片，那么 s[n:n] 的长度是多少？
	2) s[n:n+1] 的长度又是多少？
	*/
	fmt.Printf("%d \n", len(slice[3:3])) //0
	fmt.Printf("%d \n", len(slice[1:2])) //1

	slice2 := arr[0:0]
	fmt.Printf("%d \n", len(slice2)) //0

	//结论：如果 n==n 则长度为0
}
