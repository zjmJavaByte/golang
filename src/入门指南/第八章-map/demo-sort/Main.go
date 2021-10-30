package main

import (
	"fmt"
	"sort"
)

/**
对map进行排序：
如果你想为 map 排序，需要将 key（或者 value）拷贝到一个切片，再对切片排序
*/
var (
	barVal = map[string]int{
		"alpha": 34, "bravo": 56, "charlie": 23,
		"delta": 87, "echo": 56, "foxtrot": 12,
		"golf": 34, "hotel": 16, "indio": 87,
		"juliet": 65, "kili": 43, "lima": 98}
)

func main() {
	fmt.Println("未排序-----")
	for k, v := range barVal {
		fmt.Printf("Key: %v, Value: %v / ", k, v)
	}
	fmt.Println("\n排序后的---------")
	slice := make([]string, len(barVal))
	i := 0
	for k := range barVal {
		slice[i] = k
		i++
	}

	sort.Strings(slice)
	for _, s := range slice {
		fmt.Printf("Key: %v, Value: %v / ", s, barVal[s])
	}

}

//但是如果你想要一个排序的列表你最好使用结构体切片，这样会更有效  TODO
type name struct {
	value int
	key   string
}
