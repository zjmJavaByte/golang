package main

import (
	"fmt"
	"sort"
)

type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

type intArray []int

func main() {
	var data = intArray{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 758}
	//Sort(data)
	sort.Ints(data)
	for i := range data {
		fmt.Println(data[i])
	}

}

func (a intArray) Len() int {
	return len(a)
}

func (a intArray) Less(i, j int) bool {
	return a[i] > a[j]
}

func (a intArray) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func Sort(data Sorter) {
	for i := 0; i < data.Len(); i++ {
		for j := 1 + i; j < data.Len(); j++ {
			if data.Less(i, j) {
				data.Swap(i, j)
			}
		}
	}
}
