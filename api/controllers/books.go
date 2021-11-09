package controllers

import (
	"book-store/infrastructure/data/repositories"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func List(w http.ResponseWriter, r *http.Request) {
	log.Println("book - list")
	w.Header().Set("Content-Type", "application/json")
	books := repositories.List()

	json.NewEncoder(w).Encode(books)
}

func Get(w http.ResponseWriter, r *http.Request) {
	log.Printf("book - get")
	id := r.URL.Query().Get("bookId")
	w.Header().Set("Content-Type", "application/json")
	convertedId, _ := strconv.Atoi(id)
	book := repositories.Get(convertedId)
	json.NewEncoder(w).Encode(book)
}
