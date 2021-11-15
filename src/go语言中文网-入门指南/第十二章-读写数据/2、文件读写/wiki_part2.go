package main

import (
	"io/ioutil"
)

func main() {

	page := Page{"中文国际", []byte("12345678iuytrewasdfghj,mnbvcx")}
	page.save2()
	page.load2()
}

func (p *Page) save2() (err error) {
	return ioutil.WriteFile(p.Title, p.Body, 0644)
}

func (p *Page) load2() (err error) {
	p.Body, err = ioutil.ReadFile(p.Title)
	return
}
