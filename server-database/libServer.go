package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	Id     int    `json:"id"`
	Author string `json:"author"`
	Title  string `json:"title"`
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetBook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/book/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var book Book
	err = s.db.QueryRow("SELECT id, author, title FROM book WHERE id = ?", id).Scan(&book.Id, &book.Author, &book.Title)
	if err == sql.ErrNoRows {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Database error :"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (s *Store) AddBook(w http.ResponseWriter, r *http.Request) {
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

	result, err := s.db.Exec("INSERT INTO book (author, title) VALUES (?, ?)", book.Author, book.Title)
	if err != nil {
		http.Error(w, "Failed to add book. err: "+err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	book.Id = int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func (s *Store) UpdateBook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/book/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid book ID: "+err.Error(), http.StatusBadRequest)
		return
	}

	var updatedBook Book
	b, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body:"+err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(b, &updatedBook)
	if err != nil {
		http.Error(w, "Invalid request body:"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = s.db.Exec("UPDATE book SET author = ?, title = ? WHERE id = ?", updatedBook.Author, updatedBook.Title, id)
	if err != nil {
		http.Error(w, "Failed to update book: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Book with ID %d updated", id)))
}

func (s *Store) DeleteBook(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/book/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid book ID:"+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = s.db.Exec("DELETE FROM book WHERE id = ?", id)
	if err != nil {
		http.Error(w, "Failed to delete book:"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Book with ID %d deleted", id)))
}

func main() {
	//connection to the database
	db, err := sql.Open("mysql", "root:password@tcp/sample")
	if err != nil {
		fmt.Errorf("Error connecting to database %v", err)
	}
	defer db.Close()

	// testing connection using ping
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	fmt.Println("Successfully connected to the MySQL database!")

	store := NewStore(db)

	http.HandleFunc("/book/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			store.GetBook(w, r)
		case http.MethodPut:
			store.UpdateBook(w, r)
		case http.MethodDelete:
			store.DeleteBook(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			store.AddBook(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Start the server
	port := 8080
	fmt.Println("Server is running at", port)
	http.ListenAndServe(":8080", nil)
}
