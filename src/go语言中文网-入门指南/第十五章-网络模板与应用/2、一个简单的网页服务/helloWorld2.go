package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer2(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "<h1>%s<h1><div>%s</div>", "title", "body")
}

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(HelloServer2))
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
