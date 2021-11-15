package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

/**
使用闭包调试
*/
func main() {

	//准确地知道哪个文件中的具体哪个函数正在执行
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	i := len("123")
	fmt.Println(i)
	// some code
	where()
	log.Println(123)
	// some more code
	//where()

	//计算函数执行时间
	start := time.Now()
	//longCalculation()
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

}
