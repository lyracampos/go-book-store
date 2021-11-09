package routes

import (
	"book-store/api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Load() {
	router := mux.NewRouter()

	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/books", controllers.List).Methods(http.MethodGet)
	api.HandleFunc("/books/{bookId}", controllers.Get).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
