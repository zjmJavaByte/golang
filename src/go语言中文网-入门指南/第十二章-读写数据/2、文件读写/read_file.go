package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
3) 按列读取文件中的数据
*/

func main() {
	file, err := os.Open("C:\\Users\\Administrator\\Desktop\\input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1, col2, col3, col4 []string
	for {
		var v1, v2, v3, v4 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3, &v4)
		// scans until newline
		if err != nil {
			break
		}
		col1 = append(col1, v1)
		col2 = append(col2, v2)
		col3 = append(col3, v3)
		col4 = append(col4, v4)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
	fmt.Println(col4)

	filename := filepath.Base("C:\\Users\\Administrator\\Desktop\\input.txt")
	fmt.Println(filename)
}
