package main

import "fmt"

func main() {

	map1 := map[string]int{"key3": 1, "key2": 2}

	val, isPresent := map1["key1"]
	if isPresent {
		fmt.Printf("map1['key1'] 的值是：%d \n", val)
	} else {
		fmt.Println("map1['key1'] 的key不存在")
	}
	if val, ok := map1["key2"]; ok {
		fmt.Printf("map1['key2'] 的值是：%d \n", val)
	} else {
		fmt.Println("map1['key2'] 的key不存在")
	}
	delete(map1, "key2")
	if _, ok := map1["key2"]; ok {
		fmt.Printf("map1['key2'] 的值是：%d \n", val)
	} else {
		fmt.Println("map1['key2'] 的key已经被删除")
	}
}
