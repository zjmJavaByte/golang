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
		默认会调用String()方法的输出：
			fmt.Printf()
			fmt.Print()
			fmt.Println()
			fmt.Sprintf()
	*/

	/*
		自定义string()方法
	*/
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)  //%T 会给出类型的完全规格
	fmt.Printf("two1 is: %#v\n", two1) //%#v 会给出实例的完整输出，包括它的字段
	fmt.Println(two1)
	fmt.Sprintf("two1 is: %v\n", two1)

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
