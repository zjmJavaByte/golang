package main

import (
	"fmt"
	"strconv"
)

type Celsius2 float64

type Day2 int

var dayNames = [7]string{"1", "2", "3", "4", "5", "6", "7"}

type TZ2 int

const (
	UTC2 TZ2 = 1
)

var maps = map[TZ2]string{
	UTC2: "Universal Greenwich time"}

type Stack struct {
	arr    [4]string
	length int
}

func main() {

	var c Celsius2 = 23.2
	fmt.Printf("%.2f \n", c)

	var th Day2 = 3
	fmt.Println(th)

	fmt.Println(UTC2)

	s := new(Stack)
	s.Push("q")
	fmt.Println(s)
	s.Push("w")
	fmt.Println(s)
	s.Push("e")
	fmt.Println(s)
	s.Push("r")
	fmt.Println(s)

}

func (d Day2) String() string {
	return dayNames[d]
}

func (t TZ2) String() string {
	if zone, ok := maps[t]; ok {
		return zone
	}
	return " "
}

func (s Stack) pop() string {
	s.length = s.length - 1
	value := s.arr[s.length]
	return value
}

func (s *Stack) Push(value string) string {
	if s.length+1 > cap(s.arr) {
		return "full"
	}
	s.arr[s.length] = value
	s.length++
	return value
}

func (s Stack) String() string {
	str := ""
	for i, s2 := range s.arr {
		str += "[" + strconv.Itoa(i) + ":" + s2 + "] "
	}
	return str
}
