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
	/*
		switch :
			当程序满足某个条件，然后执行完里面的代码后，就不会像Java那样，没有break就会一直往下执行，go不一样；
		go中如果想继续执行，可以使用fallthrough 关键字来达到目的。
		特例：
			type-switch 不允许有 fallthrough ：报Cannot use 'fallthrough' in the type switch
	*/

	switch t := sh.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
		//fallthrough
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}

	classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)

}

func (c *Circle) Area() float32 {
	return c.radius * c.radius
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}
