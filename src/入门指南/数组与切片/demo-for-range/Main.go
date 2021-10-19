package main

import "fmt"

func main() {

	//定义
	/**
	for ix, value := range slice1 {
	    ...
	}
	*/

	var slice1 []int = make([]int, 4)

	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4

	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}
	fmt.Println("============")
	//通过slice1[ix]修改对应位置的值，
	for ix, value := range slice1 {
		slice1[ix] = 999
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}
	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}
	fmt.Println("============")
	//无法通过 value = 888  进行值得修改
	for ix, value := range slice1 {
		value = 888
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}
	for ix, value := range slice1 {
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}

	/**
	通过 _ 忽略掉返回的值
	*/
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}
	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}
	fmt.Println("如果你只需要索引，你可以忽略第二个变量") //但是无法只获取值
	for ix := range seasons {
		fmt.Printf("%d \n", ix)
	}
	// Output: 0 1 2 3

	fmt.Println("-------------练习---------------")
	/**
	练习
	*/

	items := [...]int{10, 20, 30, 40, 50}
	fmt.Println("通过值修改")
	for _, item := range items {
		item *= 2
	}
	for _, item := range items {
		fmt.Printf("%d \n", item)
	}
	fmt.Println("通过索引修改")
	for i := range items {
		items[i] *= 2
	}
	for _, item := range items {
		fmt.Printf("%d \n", item)
	}

	fmt.Println("问题 7.6 通过使用省略号操作符 ... 来实现累加方法。")
	a := sumOther(1, 1, 2, 3, 4, 5, 6, 7)
	fmt.Printf("%d \n", a)

	/**
	写一个 Sum 函数，传入参数为一个 32 位 float 数组成的数组 arrF，返回该数组的所有数字和。
	如果把数组修改为切片的话代码要做怎样的修改？如果用切片形式方法实现不同长度数组的的和呢
	*/
	fmt.Println("-----------------")

	arrF := [32]float64{4.5, 6.7, 8.9, 0.1}
	f := sum(arrF)
	fmt.Printf("%v \n", f)

	fmt.Println("-----------------")
	ints := []int{1, 2, 3, 4, 5, 6}
	f2, average := SumAndAverage(ints)
	fmt.Printf("平均值 %v  和 %v", average, f2)

}

func sumOther(other ...int) (a int) {
	for _, i2 := range other {
		a += i2
	}
	return
}

func sum(other [32]float64) (a float64) {
	for _, i2 := range other {
		a += i2
	}
	return
}

func SumAndAverage(a []int) (int, float32) {
	sum := 0
	for _, item := range a {
		sum += item
	}
	return sum, float32(sum / len(a))
}
