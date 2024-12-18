package main

import (
	"fmt"
	"net/http"
)

func HandleGreet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid", http.StatusMethodNotAllowed)
	}

	name := r.PathValue("vaidehi")
	if name == "" {
		name = "world"
	}

	fmt.Fprint(w, "Hello", name)
}

func main() {
	http.HandleFunc("/{id}", HandleGreet)
	http.ListenAndServe(":8080", nil)

}
