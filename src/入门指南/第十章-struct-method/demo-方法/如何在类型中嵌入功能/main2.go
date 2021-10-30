package main

import "fmt"

type Integer int

type Integer2 struct{ n int }

func main() {

	i := *new(Integer)
	i = 10
	fmt.Println(i.get())

	var v Integer
	v = 20
	i2 := f(v.get())
	fmt.Println(i2)

	var v2 Integer2
	v2.n = 30
	i3 := f(v2.get())
	fmt.Println(i3)
}

func (receiver Integer) get() int {
	return int(receiver)
}

func (receiver Integer2) get() int {
	return receiver.n
}

func f(i int) int {
	return i
}
