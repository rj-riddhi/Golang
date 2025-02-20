package main

import (
	"log"
	"net/http"

	"githlab.com/radhika.parmar/go-jwt-auth-project/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.JwtAuthRoutes(r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
