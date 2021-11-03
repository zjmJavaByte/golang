package main

import "fmt"

type obj interface{}

func main() {

	f := func(o obj) obj {
		switch o.(type) {
		case int:
			return o.(int) * 2 //TODO 将interface转换为相应的类型的值
		case string:
			return o.(string) + o.(string)
		}
		return o
	}

	isl := []obj{0, 1, 2, 3, 4, 5}
	objs := make([]obj, len(isl))
	for i, o := range isl {
		o2 := mapFunc2(o)
		objs[i] = o2
	}
	fmt.Println(objs)

	res2 := mapFunc(f, "0", "1", "2", "3", "4", "5")
	fmt.Println(res2)
}

func mapFunc(a func(o obj) obj, b ...obj) []obj {
	objs := make([]obj, len(b))
	for i, o := range b {
		objs[i] = a(o)
	}
	return objs
}

func mapFunc2(o obj) obj {
	switch o.(type) {
	case int:
		return o.(int) * 2 //TODO 将interface转换为相应的类型的值
	case string:
		return o.(string) + o.(string)
	}
	return o
}
