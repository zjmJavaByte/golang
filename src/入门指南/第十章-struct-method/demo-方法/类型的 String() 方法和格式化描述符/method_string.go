package main

import (
	"fmt"
	"strconv"
)

type TwoInts struct {
	a int
	b int
}

func main() {

	/*
		自定义string()方法
	*/
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)
	fmt.Println(two1.String())

	/*
		注意事项
	*/
	//t.String() 调用 fmt.Sprintf，而 fmt.Sprintf 又会反过来调用 t.String()...），很快就会导致内存溢出
	//t := new(TT)
	//t. String()
}

func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}

type TT float64

func (t TT) String() string {
	return fmt.Sprintf("%v", t)
}
