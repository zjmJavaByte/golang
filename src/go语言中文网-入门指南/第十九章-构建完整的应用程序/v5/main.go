package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

var store Store

func main() {
	store = NewUrlStore("store.gob")
	rpc.Register(store)
	rpc.HandleHTTP()
	http.HandleFunc("/add", Add)
	http.HandleFunc("/", Redirect)

	http.ListenAndServe(":8080", nil)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	var url string
	err := store.Get(&key, &url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, AddForm)
		return
	}
	var key string
	err := store.Put(&url, &key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
	fmt.Fprintf(w, "http://localhost:8080/%s", key)
}

const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`
