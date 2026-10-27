package main

import (
	"log"
	"net/http"

	"github.com/Sparsh017-byte/book-management-system/pkg/models"
	"github.com/Sparsh017-byte/book-management-system/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	if err := models.Init(); err != nil {
		log.Fatal("Database initialization failed: ", err)
	}

	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	log.Println("Server running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
