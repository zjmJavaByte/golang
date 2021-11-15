package main

import (
	"fmt"
)

func main() {

	/*os.PathError{}
	json.SyntaxError*/

	i, err := returnError()
	fmt.Println(i)
	fmt.Println(err)

}

//用 fmt 创建错误对象
func returnError() (int, error) {
	return 0, fmt.Errorf("math: square root of negative number %d", 0)
}
