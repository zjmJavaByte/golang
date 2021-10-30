package main

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

/*

make()支持三种类型：
slices  /  maps / channels（见第 14 章）
*/
func main() {
	// OK
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	//试图 make() 一个结构体变量，会引发一个编译错误
	// NOT OK
	//z := make(Bar) // 编译错误：cannot make type Bar
	//(*z).thingOne = "hello"
	//(*z).thingTwo = 1

	// OK
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	//new(Foo) 返回的是一个指向 nil 的指针，它尚未被分配内存
	// NOT OK
	u := new(Foo)
	(*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
	(*u)["y"] = "world"
}
