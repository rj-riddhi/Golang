package middleware

import (
	// "fmt"
	"encoding/json"
	"net/http"

	helper "githlab.com/radhika.parmar/go-jwt-auth-project/helpers"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientToken := r.Header.Get("Authorization")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		if clientToken == "" {
			json.NewEncoder(w).Encode(map[string]string{"error": "Please provide access token"})
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			json.NewEncoder(w).Encode(map[string]string{"error": err})
			return
		}
		w.Header().Set("email", claims.Email)
		w.Header().Set("first_name", claims.First_name)
		w.Header().Set("last_name", claims.Last_name)
		w.Header().Set("user_type", claims.User_type)
		w.Header().Set("user_id", claims.User_id)
		next.ServeHTTP(w, r)
	})

}
