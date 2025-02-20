package routes

import (
	"githlab.com/radhika.parmar/go-jwt-auth-project/controllers"
	"githlab.com/radhika.parmar/go-jwt-auth-project/middleware"
	"github.com/gorilla/mux"
)

var JwtAuthRoutes = func(router *mux.Router) {
	// Create a subrouter for secured routes
	secure := router.PathPrefix("/").Subrouter()
	secure.Use(middleware.Authenticate) // Apply the authentication middleware

	router.HandleFunc("/user/signup", controllers.UserSignup).Methods("POST")
	router.HandleFunc("/user/login", controllers.UserLogin).Methods("POST")
	secure.HandleFunc("/users", controllers.GetAllUser).Methods("GET")
	secure.HandleFunc("/users/{userId}", controllers.GetUserById).Methods("GET")
}
