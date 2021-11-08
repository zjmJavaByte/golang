package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {

	p := new(Page)
	p.Title = "中文国际"
	outputFile := "C:\\Users\\Administrator\\Desktop\\" + p.Title + ".txt"
	readFile, err := ioutil.ReadFile("C:\\Users\\Administrator\\Desktop\\input.txt")
	if err != nil {
		fmt.Printf("文件读取失败 %v", err)
	}
	p.Body = readFile
	p.save(outputFile)

	load(p.Title)
}

func (p *Page) save(outputFile string) {
	ioutil.WriteFile(outputFile, p.Body, 0644)
}

func load(title string) {
	intputFile := "C:\\Users\\Administrator\\Desktop\\" + title + ".txt"
	file, err := os.OpenFile(intputFile, os.O_RDONLY, 0)
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(file)
	defer file.Close()
	for true {
		readString, err := reader.ReadString('\n')
		fmt.Println(readString)
		if err == io.EOF {
			break
		}
	}

}
