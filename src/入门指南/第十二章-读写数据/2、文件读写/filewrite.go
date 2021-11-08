package main

import "os"

func main() {

	os.Stdout.WriteString("Hello world \n")
	f, _ := os.OpenFile("C:\\Users\\Administrator\\Desktop\\output.txt", os.O_CREATE|os.O_WRONLY, 0)

	defer f.Close()
	f.WriteString("hello, world in a file\n")
}
