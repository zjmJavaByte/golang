package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// var outputWriter *bufio.Writer
	// var outputFile *os.File
	// var outputError os.Error
	// var outputString string

	/*
		OpenFile 函数有三个参数：文件名、一个或多个标志（使用逻辑运算符 “|” 连接），使用的文件权限

		我们通常会用到以下标志：

		os.O_RDONLY：只读
		os.O_WRONLY：只写
		os.O_CREATE：创建：如果指定文件不存在，就创建该文件。
		os.O_TRUNC：截断：如果指定文件已存在，就将该文件的长度截为 0。
		在读文件的时候，文件的权限是被忽略的，所以在使用 OpenFile 时传入的第三个参数可以用 0。而在写文件时，不管是 Unix 还是 Windows，都需要使用 0666。
	*/
	outputFile, outputError := os.OpenFile("C:\\Users\\Administrator\\Desktop\\output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile.Close()

	//创建一个写入器（缓冲区）对象
	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello world!\n"

	for i := 0; i < 10; i++ {
		//将内容写入缓冲区
		outputWriter.WriteString(outputString)
	}
	//缓冲区的内容紧接着被完全写入文件
	outputWriter.Flush()

	//如果写入的东西很简单，我们可以使用 fmt.Fprintf(outputFile, “Some test data.\n”) 直接将内容写入文件
	outputFile2, outputError2 := os.OpenFile("C:\\Users\\Administrator\\Desktop\\output2.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError2 != nil {
		fmt.Printf("An error occurred with file opening or creation\n")
		return
	}
	defer outputFile2.Close()
	//fmt 包里的 F 开头的 Print 函数可以直接写入任何 io.Writer
	fmt.Fprintf(outputFile2, "Some test data.\n")
}
