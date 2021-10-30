package main

import (
	"errors"
	"fmt"
)

func main() {

	/**
	copy函数
	*/
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)
	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	/**
	append 函数
	param2 : 省略号操作符
	*/
	sl3 := []int{0, 1, 2, 3}
	ints := append(sl3, 4, 5, 6, 99, 9)
	fmt.Println(sl3)
	fmt.Println(ints)
	fmt.Println(&ints)
	fmt.Println(&sl3)
	fmt.Println(cap(sl3))
	fmt.Println(cap(ints))

	slice := []byte{'a', 'b', 'c'}
	slice = AppendByte(slice, 'e', 'f', 'g')
	fmt.Printf("%c \n", slice)

	fmt.Println("---------------练习 7.9-------")
	/*
		练习 7.9
		给定 slice s[]int 和一个 int 类型的因子 factor，扩展 s 使其长度为 len(s) * factor
	*/
	init1 := []int{1, 2, 3, 4}
	fmt.Printf("原来的值：%v 原来的长度：%d \n", init1, len(init1))
	r := resize(init1, 7)
	fmt.Printf("值：%v 长度：%d \n", r, len(r))

	fmt.Println("---------------练习 7.10-------")
	/**
	练习 7.10
	用顺序函数过滤容器：s 是前 10 个整型的切片。构造一个函数 Filter，第一个参数是 s，第二个参数是一个 fn func(int) bool，
	返回满足函数 fn 的元素切片。通过 fn 测试方法测试当整型值是偶数时的情况。
	*/
	s := []int{1, 2, 3, 4, 5, 6, -7, -8, -9, -1}
	fn := func(a int) bool {
		if a > 0 && a%2 == 0 {
			return true
		} else {
			return false
		}
	}
	filter := Filter(s, fn)
	fmt.Printf("过滤后的值：%v \n", filter)
	/*
		练习 7.11
		写一个函数 InsertStringSlice 将切片插入到另一个切片的指定位置。
	*/
	fmt.Println("---------------练习 7.11-------")
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b := []int{'q', 'w', 'e', 'r'}
	insertStringSlice, err := InsertStringSlice(&a, &b, 8)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(insertStringSlice)
	}

	fmt.Println("---------------练习 7.12-------")
	/*
		练习 7.12
		写一个函数 RemoveStringSlice 将从 start 到 end 索引的元素从切片 中移除。
	*/
	slice2 := []int{1, 2, 3, 4}
	init := slice2[0:]
	fmt.Println(init)
	stringSlice, err := RemoveStringSlice(1, 5, slice2)
	if err == nil {
		fmt.Println(stringSlice)
	} else {
		fmt.Println(err)
	}

}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n] //改变切片长度的过程称之为切片重组 reslicing
	copy(slice[m:n], data)
	return slice
}

func RemoveStringSlice(start, end int, slice []int) ([]int, error) {
	length := len(slice)
	if start > length {
		return nil, errors.New("start不在范围内")
	}
	if end > length {
		end = length - 1
	}
	n := length - (end - start + 1)
	newSlice := make([]int, n)
	copy(newSlice[0:start], slice[0:start])
	copy(newSlice[start:n], slice[end+1:])
	return newSlice, nil
}

func RemoveStringSlice2(slice []string, start, end int) []string {
	result := make([]string, len(slice)-(end-start))
	at := copy(result, slice[:start])
	copy(result[at:], slice[end:])
	return result
}

func InsertStringSlice(a, b *[]int, c int) ([]int, error) {
	ia := *a
	if c > len(ia) {
		return nil, errors.New("插入的位置不存在")
	}
	ib := *b
	i1 := len(ia)
	i2 := len(ib)
	length := i1 + i2
	ints := make([]int, length)
	copy(ints[0:c], ia[0:c])
	copy(ints[c:c+i2], ib[:])
	copy(ints[c+i2:], ia[c:])
	return ints, nil
}

func InsertStringSlice2(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}

func Filter(s []int, fn func(int) bool) []int {
	ints := make([]int, 0)
	for _, v := range s {
		if fn(v) {
			ints = append(ints, v)
		}
	}
	return ints
}

// Filter returns a new slice holding only
// the elements of s that satisfy f()
func Filter2(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, i := range s {
		if fn(i) {
			p = append(p, i)
		}
	}
	return p
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

/* [0 2 4 6 8] */

func resize(s []int, factor int) (r []int) {
	i := len(s) * factor
	ints := make([]int, i)
	copy(ints, s)
	return ints
}
