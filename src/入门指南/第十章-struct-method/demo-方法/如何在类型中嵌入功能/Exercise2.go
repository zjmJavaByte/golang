package main

import (
	"fmt"
)

type Base2 struct{}

func (Base2) Magic() {
	fmt.Println("Base2 magic")
}

func (self Base2) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base2
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func main() {

	v := new(Voodoo)
	v.Magic()     //voodoo magic
	v.MoreMagic() //Base2 magic Base2 magic
}
