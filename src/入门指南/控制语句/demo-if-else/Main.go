package main

import (
	"fmt"
	"runtime"
)

func main() {

	if true {
	}

	if true {
	} else {

	}

	if true {

	} else if true {

	}

	if runtime.GOOS == "windows" {
		//.   ..
	} else { // Unix-like
		//.   ..
	}

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	//if 可以包含一个初始化语句（如：给一个变量赋值）。这种写法具有固定的格式（在初始化语句后方必须加上分号）：
	max := 9
	val := 10
	if val > max {
		// do something
	}

	/**
	使用简短方式 := 声明的变量的作用域只存在于 if 结构中（在 if 结构的大括号之间，如果使用 if-else 结构则在 else 代码块中变量也会存在）。
	如果变量在 if 结构之前就已经存在，那么在 if 结构中，该变量原来的值会被隐藏
	*/
	if val := 10; val > max {
		// do something
	}

}

var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
	fmt.Println(prompt)
}
