package main

import (
	"io"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>Hello Go!</h1>")
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
