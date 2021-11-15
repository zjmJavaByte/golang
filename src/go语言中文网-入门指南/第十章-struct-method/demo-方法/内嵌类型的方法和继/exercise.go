package main

import "fmt"

type Car2 struct {
	wheelCount int
}

type Engine2 struct {
}

type Mercedes struct {
	Car2
}

func main() {

	m := new(Mercedes)
	m.numberOfWheels()
	fmt.Println()
}

func (receiver Car2) numberOfWheels() int {
	return receiver.wheelCount

}

func (receiver Mercedes) sayHiToMerkel() {

}
