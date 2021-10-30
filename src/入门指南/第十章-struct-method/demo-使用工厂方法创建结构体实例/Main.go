package main

import (
	"fmt"
	"unsafe"
)

/*
结构体工厂
*/

type File struct {
	fd   int    // 文件描述符
	name string // 文件名
}

type matrix struct {
	matrix string
}

func main() {

	f := NewFile(1000000, "./test.txt")
	fmt.Println(f)
	size := unsafe.Sizeof(f)
	fmt.Println(size)

}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	return &File{fd, name}
}

/*func () NewMatrix() *matrix {
	m := new(matrix) // 初始化 m
	return m
}*/
