package main

import (
	"./person"
	"fmt"
)

type s struct {
	s2   S2
	name string
}

type S2 struct {
	name2 string
}

func main() {

	p := new(person.Person)
	//p.firstName
	p.SetFirstName("z")
	name := p.GetFirstName()
	fmt.Println(name)

	/*
		结构之间的组合————》》继承
		s3 := new(s)
		s3.s2.name2 = "123"
		fmt.Println(s3)
		fmt.Println(s3.s2)*/
}
