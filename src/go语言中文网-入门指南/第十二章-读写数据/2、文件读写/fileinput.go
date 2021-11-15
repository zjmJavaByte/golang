package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
 读文件
*/

func main() {

	open, err := os.Open("C:\\Users\\Administrator\\Desktop\\input.txt")
	if err != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}

	defer open.Close()

	inputReader := bufio.NewReader(open)
	for {
		line, isPrefix, err := inputReader.ReadLine()
		inputString, _ := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s \n", inputString)
		if err == io.EOF {
			return
		}
		fmt.Printf("The input was: %s     isPrefix： %v \n", line, isPrefix)

	}

}
