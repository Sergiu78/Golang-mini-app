package main

import (
	"fmt"
	"net/http"
)

func home_page(page http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(page, "hello first page go")
}

func main() {
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":8000", nil)
}
