package main

import (
	"fmt"
	"net/http"
)

var store = NewUrlStore()

func main() {

	http.HandleFunc("/add", add)
	http.HandleFunc("/", redirect)

	http.ListenAndServe(":8080", nil)

}

func add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, AddForm)
		return
	}
	key := store.Put(url)
	fmt.Fprintf(w, "http://localhost:8080/%s", key)
	return
}

func redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := store.Get(key)
	if url == "" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

const AddForm = `

<form method="POST" action="/add">

URL: <input type="text" name="url">

<input type="submit" value="Add">

</form>

`
