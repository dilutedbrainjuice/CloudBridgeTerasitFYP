package main

import (
	"context"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
)

type ID int64
type Username string

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

		// Use the exact types directly for context keys
		var userIDKey ID = 0          // Assuming ID(0) is the zero value for your ID type
		var userNameKey Username = "" // Assuming Username("") is the zero value for your Username type

		ctx := context.WithValue(r.Context(), userIDKey, claims.ID)
		ctx = context.WithValue(ctx, userNameKey, claims.Username)

		// Logging to check the values stored in context
		log.Println("UserID stored in context:", ctx.Value(userIDKey))
		log.Println("Username stored in context:", ctx.Value(userNameKey))

		// Call the next handler with the updated request context
		next(w, r.WithContext(ctx))
		// Token is valid, call the next handler

	}
}
