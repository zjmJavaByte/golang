package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
将整个文件的内容读到一个字符串里：
*/

func main() {
	inputFile := "C:\\Users\\Administrator\\Desktop\\input.txt"
	outputFile := "C:\\Users\\Administrator\\Desktop\\input_copy.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
	println(0600)
}
