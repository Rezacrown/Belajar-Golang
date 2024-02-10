package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Contoh data / model entity
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items []Item

// function
func main() {
	router := mux.NewRouter()

	// contoh
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// cara kirim response 1
		w.Header().Set("Content-Type", "application/json")
		// datas := map[string]any{
		// 	"name": "aqua",
		// 	"item": map[string]string{
		// 		"id":          "01",
		// 		"price":       "120 USD",
		// 		"description": "lorem insumis dolor",
		// 	},
		// }

		// json.NewEncoder(w).Encode(datas)

		// cara kirim response 2
		w.Write([]byte(`
		{"message":"Hello World"},
		{"items": 
			{
				"_id": 1,
				"name": "aqua",
			},
			{
				"_id": 2,
				"name": "lee mineral",
			},
			{
				"_id": 3,
				"name": "teh kotak",
			},
		}
		`)) // output JSON

	}).Methods("GET")

	// Endpoint untuk mendapatkan semua item
	router.HandleFunc("/items", GetItems).Methods("GET")

	// Endpoint untuk menambahkan item baru
	router.HandleFunc("/items", AddItem).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}

// Handler untuk mendapatkan semua item
func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// Handler untuk menambahkan item baru
func AddItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	_ = json.NewDecoder(r.Body).Decode(&newItem)
	items = append(items, newItem)
	w.WriteHeader(http.StatusCreated)
}
