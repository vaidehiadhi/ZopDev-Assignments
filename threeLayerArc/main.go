package main

import (
	"database/sql"
	"fmt"
	"github.com/vaidehiadhi/threeLayerArc/handler"
	"github.com/vaidehiadhi/threeLayerArc/service"
	"github.com/vaidehiadhi/threeLayerArc/store"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/sample")
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error verifying connection to the database: %v", err)
	}

	store := store.NewStore(db)
	service := service.NewUserService(store)
	handler := handler.NewUserHandler(service)

	r := mux.NewRouter()

	r.HandleFunc("/user/{name}", handler.GetUser).Methods("GET")
	r.HandleFunc("/user", handler.AddUser).Methods("POST")
	r.HandleFunc("/user/{name}", handler.UpdateUser).Methods("PUT")
	r.HandleFunc("/user/{name}", handler.DeleteUser).Methods("DELETE")

	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
