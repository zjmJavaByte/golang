package main

import (
	"fmt"
	"os"
	"text/template"
)

/*
字段替代: {{.FieldName}}
*/
type Person struct {
	Name                string
	nonExportedAgeField string // 译者添加： 原文中没有定义这个，代码会报错
}

func main() {

	t := template.New("hello")
	parse, err := t.Parse("hello {{.Name}}")
	//结构体包含了一个不能导出的字段，并且当我们尝试通过一个定义字符串去合并他时
	//nonExportedAgeField is an unexported field of struct type *main.Person
	//parse, err := t.Parse("hello {{.Name}}  {{.nonExportedAgeField}}")
	if err != nil {
		fmt.Println(err)
		return
	}
	p := &Person{"朱健民", "123"}
	err = parse.Execute(os.Stdin, p)
	if err != nil {
		fmt.Println(err)
		return
	}

	t.Parse("{{.}}")
	t.Execute(os.Stdin, p)
	/*
		当模板应用在浏览器中时，要先用 html 过滤器去过滤输出的内容，像这样： {{html .}} ，或者使用一个 FieldName {{ .FieldName |html }}

		|html 部分告诉 template 引擎在输出 FieldName 的值之前要通过 html 格式化它。他会转义特殊的 html 字符（ 如：会将 > 替换成 &gt; ）, 这样可以防止用户的数据破坏 HTML 表单。

	*/

}
