package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Float64Array []float64

func main() {
	f1 := NewFloat64Array()
	f1.Fill(10)
	fmt.Printf("Before sorting %s\n", f1)
	sort.Sort(f1)
	fmt.Printf("After sorting %s\n", f1)
	if sort.IsSorted(f1) {
		fmt.Println("The float64 array is sorted!")
	} else {
		fmt.Println("The float64 array is NOT sorted!")
	}

}

func NewFloat64Array() Float64Array {
	return make([]float64, 25)
}

func (p Float64Array) String() string {
	return p.List()
}

func (p Float64Array) Fill(n int) {
	nanosecond := time.Now().Nanosecond() //获取随机数
	rand.Seed(int64(nanosecond))
	for i := 0; i < n; i++ {
		p[i] = 100 * (rand.Float64())
	}
}

func (p Float64Array) List() string {
	s := "{ "
	for i := 0; i < p.Len(); i++ {
		if p[i] == 0 {
			continue
		}
		s += fmt.Sprintf("%3.1f ", p[i])
	}
	s += " }"
	return s
}

func (p Float64Array) Len() int {
	return len(p)
}

func (p Float64Array) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p Float64Array) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
