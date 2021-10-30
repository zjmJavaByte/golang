package main

import (
	"fmt"
	"math"
	"time"
)

type Address struct {
	num int
}

type VCard struct {
	name     string
	birthday time.Time
	headImg  string
	address  map[string]*Address
}

type Point struct {
	X, Y float64
}

type Point3 struct {
	X, Y, Z float64
}

type Polar struct {
	R, T float64
}

type Rectangle struct {
	h, w int
}

func main() {
	adds := make(map[string]*Address)
	adds["now"] = &Address{123}
	adds["your"] = &Address{456}
	birthday := time.Date(1995, 10, 18, 15, 4, 5, 0, time.Local)
	vCard := &VCard{"zjm", birthday, "http", adds}
	fmt.Printf("Here is the full VCard: %v\n", vCard)

	p1 := new(Point)
	p1.X = 3
	p1.Y = 4
	fmt.Printf("The length of the vector p1 is: %f\n", Abs(p1))

	r1 := Rectangle{4, 3}
	fmt.Println("Rectangle is: ", r1)
	fmt.Println("Rectangle area is: ", area(&r1))
	fmt.Println("Rectangle area is: ", r1.Area())
}

func Abs(p *Point) (r float64) {
	f := p.X*p.X + p.Y*p.Y
	r = math.Sqrt(f)
	return
}

//注意以下两种不同的函数参数位置
//调用形式不同：area(&r1)
func area(r *Rectangle) int {
	return r.h * r.w
}

//调用形式不同：r1.Area()
func (r *Rectangle) Area() int {
	return r.h * r.w
}
