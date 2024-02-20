package main

import (
	"fmt"
	"net/http"
)

const AddForm = `
<form method="POST" action="/add">
UR1L: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

var store = NewURLStore()

func main() {
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(":8080", nil)
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path[1:]
	url := store.Get(key)
	if url == "" {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}

func Add(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		fmt.Println(w, "invalid url")
		return
	}

	key := store.Put(url)
	fmt.Fprintf(w, "http://localhost:8080/%s", key)
}
