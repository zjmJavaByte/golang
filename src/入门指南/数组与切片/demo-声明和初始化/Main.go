package main

import "fmt"

const (
	WIDTH  = 1920
	HEIGHT = 1080
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func main() {

	/**
	定义：
	1、同构的数据结构
	2、数组长度也是数组类型的一部分，所以 [5] int 和 [10] int 是属于不同类型的
	3、数组长度必须是一个常量表达式，并且必须是一个非负整数
	4、数组长度最大为 2Gb
	格式：
	var identifier [len]type  --》》 var arr1 [5]int
	var identifier new([len]type) --》》 var arr1 = new([5]int)
	*/

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
	arr33 := arr1
	fmt.Printf("%p %T \n", arr33, arr33)
	arr3[2] = 100
	//地址拷贝
	arr4 := &arr2
	arr4[2] = 100
	fmt.Printf("%v %T \n", arr2, arr2)
	//值拷贝
	arr5 := arr2
	arr5[3] = 100
	fmt.Printf("%v %T \n", arr2, arr2)

	fmt.Println("-------------练习--------------------")
	/*
		练习 7.1：array_value.go: 证明当数组赋值时，发生了数组内存拷贝。

		练习 7.2：for_array.go: 写一个循环并用下标给数组赋值（从 0 到 15）并且将数组打印在屏幕上。

		练习 7.3：fibonacci_array.go: 在第 6.6 节我们看到了一个递归计算 Fibonacci 数值的方法。但是通过数组我们可以更快的计算出 Fibonacci 数。完成该方法并打印出前 50 个 Fibonacci 数字。
	*/
	//练习 7.1：array_value.go: 证明当数组赋值时，发生了数组内存拷贝。
	var arr6 [5]int
	for i := 0; i < len(arr6); i++ {
		arr6[i] = i * 2
	}
	arr7 := arr6
	arr7[2] = 100
	for i := 0; i < len(arr6); i++ {
		fmt.Printf("Array arr6 at index %d is %d\n", i, arr6[i])
	}
	fmt.Println()
	for i := 0; i < len(arr7); i++ {
		fmt.Printf("Array arr7 at index %d is %d\n", i, arr7[i])
	}
	fmt.Println("-------")

	/**
	练习 7.3：fibonacci_array.go: 在第 6.6 节我们看到了一个递归计算 Fibonacci 数值的方法。
	但是通过数组我们可以更快的计算出 Fibonacci 数。完成该方法并打印出前 50 个 Fibonacci 数字
	*/
	var arr8 [50]int
	arr8[0] = 1
	arr8[1] = 1
	for i := 2; i < 50; i++ {
		arr8[i] = arr8[i-1] + arr8[i-2]
	}
	for i, i2 := range arr8 {
		println(i, i2)
	}

	fmt.Println("-------")

	/*
		数组定义
	*/
	//1、[10]int {1, 2, 3} : 这是一个有 10 个元素的数组，除了前三个元素外其他元素都为 0
	var arrAge = [5]int{18, 20, 15, 22, 16}
	for i, i2 := range arrAge {
		println(i, i2)
	}
	//2、... 可同样可以忽略
	var arrLazy = [...]int{5, 6, 7, 8, 22}
	for i, i2 := range arrLazy {
		println(i, i2)
	}
	//3、只有索引 3 和 4 被赋予实际的值，其他元素都被设置为空的字符串
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	for i, i2 := range arrKeyValue {
		println(i, i2)
	}

	//取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
	var ss *[5]string
	ss = &arrKeyValue
	fmt.Println(ss) //&[   Chris Ron]
	//符号 * 可以放在一个指针前,那么它将得到这个指针指向地址上所存储的值
	fmt.Println(*ss) //[   Chris Ron]

	var i1 = 6
	var intP *int
	intP = &i1
	fmt.Println(intP)        //0xc000128058
	fmt.Println(*intP)       //6
	fmt.Println(&i1)         //0xc000128058
	fmt.Println(intP == &i1) //true

	for i := 0; i < 1; i++ {
		fp(&[3]int{i, i * i, i * i * i})
	}

	fmt.Println("7.1.3 多维数组")
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}

	fmt.Println("7.1.4 将数组传递给函数")
	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum(&array)  // Note the explicit address-of operator
	x2 := Sum(&array) // Note the explicit address-of operator
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f \n", x)
	fmt.Printf("The sum of the array is: %f \n", x2)

	arr9 := [1]int{1}
	fmt.Printf("%p \n", &arr9) ///0xc00000a0f8   地址拷贝
	printAddress(arr9)         //0xc00000a100     数组拷贝
	printAddress2(&arr9)       //0xc00000a0f8   地址拷贝
}

func printAddress(a [1]int) {
	fmt.Printf("%p \n", &a)
}
func printAddress2(a *[1]int) {
	fmt.Printf("%p \n", a)
}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a { // derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}

func Sum2(a *[3]float64) (sum float64) {
	for _, v := range a { // derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}

func fp(a *[3]int) {
	fmt.Println(a)         //&[0 0 0]
	fmt.Printf("%p \n", a) //0xc000010120
	fmt.Println(*a)        //[0 0 0]
}
