package main

import "fmt"

func main() {

	//数组长度最大为 2Gb

	a := [...]string{"a", "b", "c", "d"}
	slice := []int{7, 9, 3, 5, 1}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	for i := range slice {
		fmt.Println("Array item", i, "is", slice[i])
	}
	fmt.Printf("%T \n", a)
	fmt.Printf("%p \n", slice)

	fmt.Println()

	//创建方式
	var arr1 = new([5]int)
	var arr2 [5]int
	fmt.Printf("%p %T \n", arr1, arr1) //0xc00000c360 *[5]int
	fmt.Printf("%p %T \n", arr2, arr2) //%!p([5]int=[0 0 0 0 0]) [5]int
	fmt.Println(*arr1)                 //[0 0 0 0 0]  符号 * 可以放在一个指针前，如 *intP，那么它将得到这个指针指向地址上所存储的值
	//进行值得拷贝
	arr3 := *arr1
	arr3[2] = 100
	//地址拷贝
	arr4 := &arr2
	arr4[2] = 100
	fmt.Printf("%v %T \n", arr2, arr2)
}
