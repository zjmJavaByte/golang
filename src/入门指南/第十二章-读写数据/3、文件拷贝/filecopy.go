// filecopy.go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CopyFile2("C:\\Users\\Administrator\\Desktop\\output.txt", "C:\\Users\\Administrator\\Desktop\\input.txt")
	fmt.Println("Copy done!")
}

func CopyFile2(des, src string) error {
	open, err := os.Open(src)
	if err != nil {
		return err
	}
	defer open.Close()

	open2, err := os.OpenFile(des, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer open2.Close()

	_, err = io.Copy(open2, open)
	return err
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
