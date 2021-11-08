package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

var content string

/*
讲一个结构体读取到 gob文件中
*/

func main() {

	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}

	file, err := os.OpenFile("C:\\Users\\Administrator\\Desktop\\input.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	if err != nil {
		fmt.Printf("打开文件失败：%v", err)
	}
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
}
