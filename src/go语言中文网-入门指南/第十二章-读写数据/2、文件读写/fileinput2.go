package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	open, err := os.Open("C:\\Users\\Administrator\\Desktop\\input.txt")
	if err != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}

	defer open.Close()
	buf := make([]byte, 1024)
	inputReader := bufio.NewReader(open)
	for {
		read, _ := inputReader.Read(buf)
		fmt.Printf("The input was: %d \n", read)
		if read == 0 {
			break
		}

	}
}
