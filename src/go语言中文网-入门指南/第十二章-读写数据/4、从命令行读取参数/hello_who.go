package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	who := ""
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println(who)

}
