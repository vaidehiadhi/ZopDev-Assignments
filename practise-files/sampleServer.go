package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "this is a sample server")
	})

	port := 8080

	fmt.Println("the server is running at", port)
	http.ListenAndServe(":8080", nil)
}
