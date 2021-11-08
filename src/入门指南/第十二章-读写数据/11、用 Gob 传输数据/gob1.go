package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

/*
以字节缓冲模拟网络传输
*/

func main() {

	var netWork bytes.Buffer
	newEncoder := gob.NewEncoder(&netWork)
	decoder := gob.NewDecoder(&netWork)
	p := P{7, 5, 8, "朱健民"}
	err := newEncoder.Encode(&p)
	if err != nil {
		fmt.Println(err)
	}

	var q Q
	err = decoder.Decode(&q)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)
}
