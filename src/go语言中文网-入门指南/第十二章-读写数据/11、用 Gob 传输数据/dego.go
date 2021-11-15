package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"os"
)

type Address2 struct {
	Type    string
	City    string
	Country string
}

type VCard2 struct {
	FirstName string
	LastName  string
	Addresses []*Address2
	Remark    string
}

func main() {

	//var pa Address
	//var wa Address
	var vc VCard2

	open, err := os.Open("C:\\Users\\Administrator\\Desktop\\input.gob")
	defer open.Close()
	if err != nil {
		fmt.Printf("打开文件失败：%v", err)
	}
	reader := bufio.NewReader(open)
	decoder := gob.NewDecoder(reader)
	err = decoder.Decode(&vc)
	if err != nil {
		fmt.Printf("读取错误：%v", err)
	}
	fmt.Println(vc)
	for _, address := range vc.Addresses {
		fmt.Println(address)
	}
}
