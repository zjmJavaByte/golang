package main

func main() {

	/*
		1、new和make都不是go的关键字，而是go预定义的函数
		2、make返回一个变量而new返回一个变量的指针
		3、new 会通过zeroed value 來初始化型別，也就是字串會是""，number 會是 0，bool初始化为 false,channel, func, map, slice 等等則會是 nil
	*/

	/*
		1、new 和 make 都用于分配内存；
		2. new 和 make 都是在堆上分配内存；
		3. new 对指针类型分配内存，返回值是分配类型的指针，new不能直接对 slice 、map、channel 分配内存；
		4. make 仅用于 slice、map和 channel 的初始化，返回值为类型本身，而不是指针；
	*/

	//make()  func make(t Type, size ...IntegerType) Type
	//new()  func new(Type) *Type
}
