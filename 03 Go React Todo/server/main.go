package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/radhika.parmar/go-react-todo/routes"
)

func main() {
	r := mux.NewRouter()
	routes.TodoRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
