package main

import "fmt"

/*
空接口的使用：类似于Java的object
使用空接口数组接受不同的类型
*/

type Elements interface {
}

type Vector struct {
	a []Elements
}

func main() {

	vector := new(Vector)
	vector.a = make([]Elements, 4)
	vector.Set(0, 1)   //int
	vector.Set(1, "2") //string
	vector.Set(2, 3.5) //float64
	vector.Set(3, '6') //int32

	var a = '6'
	fmt.Printf("%T \n", a)

	f := func(any interface{}) {
		switch v := any.(type) {
		case bool:
			fmt.Printf("any %v is a bool type \n", v)
		case int:
			fmt.Printf("any %v is an int type \n", v)
		case float32:
			fmt.Printf("any %v is a float32 type \n", v)
		case string:
			fmt.Printf("any %v is a string type \n", v)
		case int32:
			fmt.Printf("any %v is a int32 type \n", v)
		case float64:
			fmt.Printf("any %v is a float64 type\n", v)
		default:
			fmt.Println("unknown type!")
		}
	}

	for _, item := range vector.a {
		f(item)
	}
}

func (v *Vector) At(i int) Elements {
	return v.a[i]
}

func (v *Vector) Set(i int, e Elements) {
	v.a[i] = e
}
