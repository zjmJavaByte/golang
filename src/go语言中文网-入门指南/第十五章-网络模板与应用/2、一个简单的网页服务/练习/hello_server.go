package main

import (
	"fmt"
	"net/http"
)

type hello struct {
}

func main() {
	var h hello

	http.ListenAndServe(":8080", h)

}

func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello word")
}
