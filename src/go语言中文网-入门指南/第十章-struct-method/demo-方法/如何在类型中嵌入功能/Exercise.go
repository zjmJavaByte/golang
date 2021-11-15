package main

import (
	"fmt"
)

type Base struct {
	//sync.RWMutex
	id int
}

type Person struct {
	Base
	FirstName string
	LastName  string
}

type Employee struct {
	Person
	salary float64
}

func main() {

	employee := &Employee{Person{Base{10}, "", ""}, 1.3}
	fmt.Println(employee.Id())
	employee.SetId(20)
	fmt.Println(employee.Id())
}

func (receiver *Base) Id() int {
	//receiver.RLock()
	//defer receiver.RUnlock()
	return receiver.id
}

func (receiver *Base) SetId(id int) {
	//receiver.Lock()
	//defer receiver.Unlock()
	receiver.id = id
}
