package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/Get", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			//write header and body here
			fmt.Fprintf(w, "Get request: hello")
		} else {
			http.Error(w, "invalid", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			//write header and body here
			fmt.Fprintf(w, "Post request")
		} else {
			http.Error(w, "invalid", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/put", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			fmt.Fprintf(w, "post request")
		} else {
			http.Error(w, "invalid", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			fmt.Fprintf(w, "deleted")
		} else {
			http.Error(w, "error", http.StatusMethodNotAllowed)
		}
	})

	port := 8080

	fmt.Println("server starting at: ", port)
	http.ListenAndServe(":8080", nil)

}
