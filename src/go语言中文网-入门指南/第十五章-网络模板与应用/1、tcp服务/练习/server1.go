package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var mapUser map[string]int

func main() {
	var listen net.Listener
	var err error
	var coon net.Conn
	mapUser = make(map[string]int)

	listen, err = net.Listen("tcp", "localhost:50000")
	CheckError(err)
	for true {
		coon, err = listen.Accept()
		CheckError(err)
		go doSomething(coon)
	}
}

func doSomething(conn net.Conn) {
	var buf []byte
	var error error
	for true {
		buf = make([]byte, 512)
		_, error = conn.Read(buf)
		CheckError(error)
		input := string(buf)
		if strings.Contains(input, "sh") {
			os.Exit(0)
		}
		if strings.Contains(input, "who") {
			DisplayList()
		}
		index := strings.Index(input, "say")
		s := input[:index-1]
		mapUser[s] = 1
		fmt.Printf("Received data: --%v--", string(buf))
	}
}

func CheckError(err error) {
	if err != nil {
		panic("程序错误即将退出,错误原因：" + err.Error())
	}
}

func DisplayList() {
	fmt.Println("--------------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for key, value := range mapUser {
		fmt.Printf("User %s is %d\n", key, value)
	}
	fmt.Println("--------------------------------------------")
}
