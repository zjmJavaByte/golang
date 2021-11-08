package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Title    string
	Price    float64
	Quantity int
}

func main() {
	open, err := os.Open("C:\\Users\\Administrator\\Desktop\\products.txt")
	if err != nil {
		fmt.Println("文件读取失败")
		return
	}

	defer open.Close()

	reader := bufio.NewReader(open)
	books := make([]*Book, 0)
	for {

		line, err := reader.ReadString('\n')
		readErr := err
		/*if err != nil {
			fmt.Println("读取错误")
		}*/
		split := strings.Split(line, ";")

		b := new(Book)

		b.Title = split[0]
		b.Price, err = strconv.ParseFloat(split[1], 32)
		if err != nil {
			fmt.Printf("Error in file: %v", err)
		}
		//strconv.ParseInt(split[2],0,0)
		b.Quantity, err = strconv.Atoi(split[2])
		books = append(books, b)

		if readErr == io.EOF {
			break
		}
	}

	fmt.Println("We have read the following books from the file: ")
	for _, bk := range books {
		fmt.Println(bk)
	}
}
