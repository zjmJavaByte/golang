// json.go
package main

/*
序列化
*/

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:

	//在 web 应用中最好使用 json.MarshalforHTML() 函数，其对数据执行 HTML 转码，所以文本可以被安全地嵌在 HTML <script> 标签中
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s \n", js)

	var v VCard
	json.Unmarshal(js, &v)
	fmt.Printf("反序列化：%s \n", v)
	file, _ := os.OpenFile("C:\\Users\\Administrator\\Desktop\\input.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()

	/*
		要想把 JSON 直接写入文件，可以使用 json.NewEncoder 初始化文件（或者任何实现 io.Writer 的类型），
		并调用 Encode ()；反过来与其对应的是使用 json.Decoder 和 Decode () 函数：
	*/
	encoder := json.NewEncoder(file)

	err := encoder.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}

	/*
		不是所有的数据都可以编码为 JSON 类型：只有验证通过的数据结构才能被编码：

		JSON 对象只支持字符串类型的 key；要编码一个 Go map 类型，map 必须是 map [string] T（T 是 json 包中支持的任何类型）
		Channel，复杂类型和函数类型不能被编码
		不支持循环数据结构；它将引起序列化进入一个无限循环
		指针可以被编码，实际上是对指针指向的值进行编码（或者指针是 nil）
	*/
}
