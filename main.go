package main

import (
	"io"
	"net/http"
)

func main() {
	// this is a comment
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "<h1>hello, bob!</h1>\n")
	})
	http.ListenAndServe(":5000", nil)
}
