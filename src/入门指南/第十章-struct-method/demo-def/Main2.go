package main

import (
	"fmt"
	"strings"
)

type Person1 struct {
	firstName string
	lastName  string
}

func upPerson(p *Person1) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func upPerson2(p Person1) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 1-struct as a value type:
	var pers1 Person1
	pers1.firstName = "Chris"
	pers1.lastName = "Woodward"
	upPerson(&pers1)
	//upPerson2(pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	// 2—struct as a pointer:
	pers2 := new(Person1)
	pers2.firstName = "Chris"
	pers2.lastName = "Woodward"
	(*pers2).lastName = "Woodward" // 这是合法的
	upPerson(pers2)
	//upPerson2(*pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	//&struct1{a, b, c} 是一种简写，底层仍然会调用 new ()
	// 3—struct as a literal:
	pers3 := &Person1{"Chris", "Woodward"}
	upPerson(pers3)
	//upPerson2(*pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}
