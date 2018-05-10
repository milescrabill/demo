package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "<h1>hello, travis!</h1>\n")
	})
	http.ListenAndServe(":5000", nil)
}
