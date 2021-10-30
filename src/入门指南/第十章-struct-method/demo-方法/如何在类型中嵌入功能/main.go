package main

import "fmt"

type Log struct {
	name string
}

type Customer struct {
	name string
	log  *Log
	Log
}

func main() {

}

func (receiver Log) print() {
	fmt.Println(receiver.name)
}
