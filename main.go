package main

import (
	"net/http"
	"io"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "<h1>hello, world!</h1>\n")
	})
	http.ListenAndServe(":5000", nil)
}