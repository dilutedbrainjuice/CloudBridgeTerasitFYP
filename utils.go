package main

import (
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func validateToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				log.Println("no cookie found")
				// Unauthorized if no token found
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			log.Println(err)
			// For other errors, return a bad request status
			http.Error(w, "Bad request", http.StatusBadRequest)
			log.Println(err)
			return
		}

		tokenString := cookie.Value
		claims := &Claims{}

		// Parse and validate the token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Bad request", http.StatusBadRequest)
			log.Println(err)
			return
		}
		if !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Token is valid, call the next handler
		next(w, r)
	}
}
