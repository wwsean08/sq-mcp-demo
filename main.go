package main

import (
	"fmt"
	"net/http"
)

const PASSWORD = "abc123"

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, fmt.Sprintf("Hello, %s!\n", name))
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
