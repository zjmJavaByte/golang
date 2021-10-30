package main

import "fmt"

func main() {

	/*
		定义
		var map1 map[keytype]valuetype

		key：可以是 string、int、float 、只包含内建类型的 struct 、指针和接口类型 、 函数
			 不可以是 切片 、 结构体
		     如果要用结构体作为 key 可以提供 Key() 和 Hash() 方法，这样可以通过结构体的域计算出唯一的数字或者字符串的 key

		key为函数：
		mf := map[int]func() int{
		        1: func() int { return 10 },
		        2: func() int { return 20 },
		        5: func() int { return 50 },
		    }

		value：可以是任意类型的
	*/
	var map1 = map[string]int{}
	var map2 = map[string]int{"key1": 1, "key2": 2}
	var map3 = make(map[string]int)

	map1["key1"] = 1
	map1["key2"] = 2

	map3["key1"] = 1
	map3["key2"] = 2

	fmt.Printf("map1 literal at \"one\" is: %d\n", map1["key1"])
	fmt.Printf("map1 literal at \"one\" is: %d\n", map1["key2"])

	fmt.Printf("map2 literal at \"one\" is: %d\n", map2["key1"])
	fmt.Printf("map2 literal at \"one\" is: %d\n", map2["key2"])

	fmt.Printf("map3 literal at \"one\" is: %d\n", map3["key1"])
	fmt.Printf("map3 literal at \"one\" is: %d\n", map3["key2"])

	//注意事项：不要使用 new，永远用 make 来构造 map，如果使用会获得一个空引用的指针
	//map4 := new(map[string]int)
	//map4["key1"] = 1    Invalid operation: 'map4["key1"]' (type '*map[string]int' does not support indexing

	/*
		定义map的容量
		make(map[keytype]valuetype, cap)
	*/
	map5 := make(map[string]float32, 100)
	i := len(map5)
	fmt.Printf("map5的容量：%d \n", i)
	map5["key1"] = 1
	i = len(map5)
	fmt.Printf("map5的容量：%d \n", i)

	/*
		用切片作为 map 的值
	*/
	//mp1 := make(map[int][]int)
	//mp2 := make(map[int]*[]int)

}
