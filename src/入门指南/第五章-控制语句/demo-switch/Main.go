package main

import "fmt"

func main() {

	a := 1

	switch a {
	case 1:
		fmt.Println("It's equal to 1")
	case 2, 3:
		fmt.Println("It's equal to 2,3")
	default:
		fmt.Println("It's not equal to 1,2,3")
	}

	/*

	 */
	//一旦成功地匹配到某个分支，在执行完相应代码后就会退出整个 switch 代码块，也就是说您不需要特别使用 break 语句来表示结束。
	var b = 0
	switch a {
	case 1:
		b = 1
	case 2:
		b = 2
	default:
		b = 13
	}
	fmt.Println(b)

	//如果在执行完每个分支的代码后，还希望继续执行后续分支的代码，可以使用 fallthrough 关键字来达到目的。
	switch a {
	case 1:
		b = 1
		fallthrough
	case 2:
		b = 2
		fallthrough
	default:
		b = 13
	}
	fmt.Println(b)

}
