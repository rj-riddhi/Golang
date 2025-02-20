package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/radhika.parmar/go-bookstore/pkg/routes"
)

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Register routes
	routes.RegisterBookStores(r)

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", r))
}
