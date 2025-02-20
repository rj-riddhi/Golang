package routes

import (
	"github.com/gorilla/mux"
	"github.com/radhika.parmar/go-react-todo/controllers"
)

var TodoRoutes = func(routes *mux.Router) {
	routes.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	routes.HandleFunc("/todo", controllers.CreateTodo).Methods("POST")
	routes.HandleFunc("/todo/{todoId}", controllers.UpdateTodo).Methods("PUT")
	routes.HandleFunc("/todo/{todoId}", controllers.DeleteTodo).Methods("DELETE")
}
