package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

type car struct {
	make  string
	model string
	price float32
}

type valuable interface {
	getValue() float32
}

func main() {

	position := stockPosition{"1", 2.3, 3.4}
	showValue(position)

	c := car{"1", "2", 3.5}
	showValue(c)

	var r io.Reader
	r = os.Stdin // see 12.1
	r = bufio.NewReader(r)
	r = new(bytes.Buffer)
	f, _ := os.Open("test.txt")
	r = bufio.NewReader(f)
}

func (c car) getValue() float32 {
	return c.price
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

func showValue(v valuable) {
	fmt.Printf("Value of the asset is %f\n", v.getValue())
}
