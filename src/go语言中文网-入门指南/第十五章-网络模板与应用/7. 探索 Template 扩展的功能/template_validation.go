package main

import (
	"log"
	"text/template"

	"fmt"
)

/*
模板验证
*/
func main() {

	defer func() {
		if x := recover(); x != nil {

			log.Printf("run time panic: %v", x)

		}

	}()

	tOk := template.New("ok")

	// 一个有效的模板，所以 Must 时候不会出现恐慌（panic）

	template.Must(tOk.Parse("/*这是一个注释 */ some static text: {{ .Name }}"))

	fmt.Println("The first one parsed OK.")

	fmt.Println("The next one ought to fail.")

	tErr := template.New("error_template")

	template.Must(tErr.Parse("some static text {{ .Name }"))

}
