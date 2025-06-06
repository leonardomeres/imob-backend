package main

import (
	"imob/internal/api"
	"imob/internal/database"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

func main() {
	db, err := database.OpenDbConnection()
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	handler := api.NewHandler(db)
	if err != nil {
		log.Fatalf("Failed to create API handler: %v", err)
	}
	http.HandleFunc("/api/customers", handler.HandleCustomers)

	// CORS for frontend access
	api := cors.Default().Handler(http.DefaultServeMux)
	log.Println("Backend listening on http://localhost:8080")
	http.ListenAndServe(":8080", api)
}
