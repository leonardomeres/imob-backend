package api

import (
	"database/sql"
	"encoding/json"
	"imob/internal/types"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) HandleCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		rows, err := h.db.Query("SELECT id, name, phone, address, listing_link, notes FROM customers")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var customers []types.Customer
		for rows.Next() {
			var c types.Customer
			err = rows.Scan(&c.ID, &c.Name, &c.Phone, &c.Address, &c.ListingLink, &c.Notes)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			customers = append(customers, c)
		}
		json.NewEncoder(w).Encode(customers)

	case "POST":
		var c types.Customer
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err := h.db.Exec("INSERT INTO customers (name, phone, address, listing_link, notes) VALUES (?, ?, ?, ?, ?)",
			c.Name, c.Phone, c.Address, c.ListingLink, c.Notes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"status": "created"})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
