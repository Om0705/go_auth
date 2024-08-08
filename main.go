package main

import (
	"go-auth-api/config"
	"go-auth-api/handlers"
	"log"
	"net/http"
)

func main() {
	db := config.InitDB()
	defer db.Close()
	http.HandleFunc("/login", handlers.LoginHandler(db))
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
