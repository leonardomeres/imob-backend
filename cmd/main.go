package main

import (
	"imob/internal/api"
	"imob/internal/database"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := database.OpenDbConnection()
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	handler := api.NewHandler(db)

	http.HandleFunc("/api/customers", handler.HandleCustomers)

	// Use your custom CORS middleware instead of cors.Default()
	withCORS := api.WithCORS(http.DefaultServeMux)

	log.Println("Backend listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", withCORS); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
