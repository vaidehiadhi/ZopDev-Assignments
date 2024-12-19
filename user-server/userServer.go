package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"regexp"

	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Phone int    `json:"phone"`
	Email string `json:"email"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name cannot be empty")
	}

	if u.Age < 0 {
		return errors.New("age cannot be negative")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailMatch, _ := regexp.MatchString(emailRegex, u.Email)

	if !emailMatch {
		return errors.New("invalid email format")
	}

	phoneRegex := `^\d{9}$`
	phoneMatch, _ := regexp.MatchString(phoneRegex, fmt.Sprintf("%d", u.Phone))
	if !phoneMatch {
		return errors.New("invalid phone number format")
	}

	return nil
}

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// func ValidateDeets(email string, phone int) error {
//	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
//	emailMatch, _ := regexp.MatchString(emailRegex, email)
//	if !emailMatch {
//		return errors.New("invalid email format")
//	}
//	phoneRegex := `^\d{10}$`
//	phoneMatch, _ := regexp.MatchString(phoneRegex, fmt.Sprintf("%d", phone))
//	if !phoneMatch {
//		return errors.New("invalid phone number format")
//	}
//
//	return nil
// }

func (s *Store) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if name == "" {
		http.Error(w, "invalid name in path", http.StatusBadRequest)
		return
	}

	var user User
	err := s.db.QueryRow("SELECT name, age, phone, email FROM `user` WHERE name = ?", name).
		Scan(&user.Name, &user.Age, &user.Phone, &user.Email)

	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "error encoding response", http.StatusInternalServerError)
	}
}

func (s *Store) AddUser(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	var user User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := user.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = s.db.Exec("INSERT INTO `user` (name, age, phone, email) VALUES (?, ?, ?, ?)", user.Name, user.Age, user.Phone, user.Email)
	if err != nil {
		http.Error(w, "failed to add user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (s *Store) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if name == "" {
		http.Error(w, "invalid name in path", http.StatusBadRequest)
		return
	}

	var updatedUser User

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request body: %v", err), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	err = json.Unmarshal(body, &updatedUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid request body: %v", err), http.StatusBadRequest)
		return
	}

	if err := updatedUser.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = s.db.Exec("UPDATE `user` SET age = ?, phone = ?, email = ? WHERE name = ?",
		updatedUser.Age, updatedUser.Phone, updatedUser.Email, name,
	)

	if err != nil {
		http.Error(w, fmt.Sprintf("failed to update user: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("user with name '%s' updated successfully", name),
	})
}

func (s *Store) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	if name == "" {
		http.Error(w, "Invalid name in path", http.StatusBadRequest)
		return
	}

	result, err := s.db.Exec("DELETE FROM `user` WHERE name = ?", name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("error checking rows affected: %v", err), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": fmt.Sprintf("user with name '%s' deleted successfully", name),
	})
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp/sample")
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("error verifying connection to the database: %v", err)
	}

	store := NewStore(db)

	r := mux.NewRouter()
	r.HandleFunc("/user/{name}", store.GetUser).Methods("GET")
	r.HandleFunc("/user", store.AddUser).Methods("POST")
	r.HandleFunc("/user/{name}", store.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{name}", store.DeleteUser).Methods("DELETE")

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
