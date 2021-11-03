package main

import "fmt"

/*
构建通用类型或包含不同类型变量的数组
*/

type Miner interface {
	Len() int
	ElemIx(ix int) interface{}
	Less(i, j int) bool
	Swap(i, j int)
}

func main() {

	array := IntArray{1, 5, 2, 3, 7, 8, 9, 2, 1, 3}
	min := Min(array)
	fmt.Println(min)

	array2 := StringArray{"z", "j", "m"}
	min2 := Min(array2)
	fmt.Println(min2)
}

func Min(miner Miner) interface{} {
	min := miner.ElemIx(0)
	for i := 1; i < miner.Len(); i++ {
		if miner.Less(i, i-1) {
			min = miner.ElemIx(i)
		} else {
			miner.Swap(i, i-1)
		}
	}
	return min
}

type IntArray []int

func (p IntArray) Len() int                  { return len(p) }
func (p IntArray) ElemIx(ix int) interface{} { return p[ix] }
func (p IntArray) Less(i, j int) bool        { return p[i] < p[j] }
func (p IntArray) Swap(i, j int)             { p[i], p[j] = p[j], p[i] }

type StringArray []string

func (p StringArray) Len() int                  { return len(p) }
func (p StringArray) ElemIx(ix int) interface{} { return p[ix] }
func (p StringArray) Less(i, j int) bool        { return p[i] < p[j] }
func (p StringArray) Swap(i, j int)             { p[i], p[j] = p[j], p[i] }
