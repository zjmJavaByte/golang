package main

import (
	"fmt"
	"sort"
)

type Person struct {
	firstName, lastName string
}

type Persons []Person

func main() {
	p1 := Person{"ss", "ss"}
	p2 := Person{"y", "ap"}
	p3 := Person{"z", "jm"}
	arrP := Persons{p1, p2, p3}
	fmt.Printf("Before sorting: %v\n", arrP)
	sort.Sort(arrP)
	fmt.Printf("After sorting: %v\n", arrP)

}

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Less(i, j int) bool {
	in := p[i].lastName + " " + p[i].firstName
	jn := p[j].lastName + " " + p[j].firstName
	return in < jn
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
