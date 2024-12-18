package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Book struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

var Books = make(map[int]Book)
var NextID = 1

// Get Request
func GetBook(w http.ResponseWriter, r *http.Request) {
	//idStr := r.PathValue("id")
	idStr := strings.TrimPrefix(r.URL.Path, "/book/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book, exists := Books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	bookJSON, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Failed to marshal book", http.StatusInternalServerError)
		return
	}
	w.Write(bookJSON)
}

// Post Request
func AddBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &book)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	book.Id = NextID
	Books[NextID] = book
	NextID++

	w.Header().Set("Content-Type", "application/json")
	bookJSON, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Failed to marshal book", http.StatusInternalServerError)
		return
	}
	w.Write(bookJSON)
}

// Put request
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	//idStr := r.PathValue("id")
	idStr := strings.TrimPrefix(r.URL.Path, "/book/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var updatedBook Book

	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &updatedBook)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	book, exists := Books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	if updatedBook.Title != "" {
		book.Title = updatedBook.Title
	}
	if updatedBook.Author != "" {
		book.Author = updatedBook.Author
	}

	Books[id] = book

	w.Header().Set("Content-Type", "application/json")
	updatedBookJSON, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Failed to marshal updated book", http.StatusInternalServerError)
		return
	}
	w.Write(updatedBookJSON)
}

// Delete Request
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//idStr := r.PathValue("id")
	idStr := strings.TrimPrefix(r.URL.Path, "/book/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	_, exists := Books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	delete(Books, id)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Book with ID %d deleted", id)))
}

func main() {
	http.HandleFunc("/book/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetBook(w, r)
		case http.MethodPut:
			UpdateBook(w, r)
		case http.MethodDelete:
			DeleteBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			AddBook(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the server
	port := 8080
	fmt.Println("Server is running at", port)
	http.ListenAndServe(":8080", nil)
}
