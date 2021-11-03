package main

import "fmt"

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func main() {

	var sh Shaper
	/*s := Square{5.6}
	sh = &s*/
	s := new(Square)
	s.side = 5.6
	sh = s
	if t, ok := sh.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := sh.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}

	/*
		注意事项：
			如果忽略 sh.(*Square) 中的 * 号，会导致编译错误：impossible type assertion:
		Square does not implement Shaper (Area method has pointer receiver)


	*/

}

func (c *Circle) Area() float32 {
	return c.radius * c.radius
}

func (s *Square) Area() float32 {
	return s.side * s.side
}
