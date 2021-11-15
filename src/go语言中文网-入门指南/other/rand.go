package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*

https://blog.csdn.net/aslackers/article/details/78548738

生成随机数
*/

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(time.Now().Nanosecond()))
	}

}
