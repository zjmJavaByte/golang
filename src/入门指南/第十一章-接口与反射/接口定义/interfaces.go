package main

import "fmt"

type Shaper interface {
	Area() float32
	Perimeter() float32
}

type Square struct {
	side float32
}

type Triangle struct {
	a, b, c float32
}

func main() {

	/*
		将接口变量赋值给结构变量
	*/
	sq1 := new(Square)
	sq1.side = 5
	var areaIntf Shaper
	areaIntf = sq1 //类似于面向对象中的接口，接口定义的方法，实现类必须都实现
	// shorter,without separate declaration:
	// areaIntf := Shaper(sq1)
	// or even:
	// areaIntf := sq1
	fmt.Printf("%p \n", areaIntf)
	fmt.Printf("%T \n", areaIntf)
	fmt.Printf("The square has area: %f\n", areaIntf.Area())

	t := new(Triangle)
	t.a = 1
	t.b = 1
	t.c = 1
	areaIntf = t //类似于面向对象中的接口，接口定义的方法，实现类必须都实现
	fmt.Printf("%p \n", areaIntf)
	fmt.Printf("%T \n", areaIntf)
	fmt.Printf("The triangle has area: %f\n", areaIntf.Area())

	/*
		接口变量会根据相应的赋值处理，找到对应的实现结构的方法
	*/
	square := &Square{1.2}
	t2 := Triangle{1, 2, 3}
	shapes := []Shaper{square, t2}
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}

}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (t Triangle) Area() float32 {
	return (t.a + t.b + t.c) / 2
}

func (sq *Square) Perimeter() float32 {
	return sq.side * sq.side
}

func (t Triangle) Perimeter() float32 {
	return (t.a + t.b + t.c) / 2
}
