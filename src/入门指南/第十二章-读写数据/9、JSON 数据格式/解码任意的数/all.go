package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)

	var i interface{}

	json.Unmarshal(b, &i)

	fmt.Println(i)
	fmt.Printf("%T \n", i)

	//m是一个map集合
	m, ok := i.(map[string]interface{})
	fmt.Println(ok)
	fmt.Println(m)
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)

		case []interface{}:
			fmt.Println(k, "is an interface array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case []string:
			fmt.Println(k, "is an string array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don’t know how to handle")
		}

	}

}
