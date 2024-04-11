package main

import (
	"database/sql"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func ConnectPostgresDB() *sql.DB {
	//database

	connStr := "user=postgres password=admin@123 dbname=cloudbridge sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
	}
	fmt.Println("Successfully connected to cloudbridge database")

	return db
}

// int for PK
func insertUser(db *sql.DB, user User) int {
	query := `INSERT INTO "user" (username, isprovider, email, password, profilepicurl, city, pc_specs, description, cloud_service, createdat) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING "ID"`

	var pk int
	err := db.QueryRow(query, user.Username, user.IsProvider, user.Email, user.Password, user.ProfilePicURL, user.City, user.PCSpecs, user.Description, user.CloudService, user.CreatedAt).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk

}

func authenticateUser(db *sql.DB, username, password string) *User {
	// Retrieve user from database based on username
	user := getUserByUsername(db, username)
	if user == nil {
		log.Println("here user not found")
		return nil // User not found
	}

	// Compare hashed password with provided password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Println(err)
		log.Println("here password not match")
		return nil // Passwords don't match
	}

	return user // Authentication successful
}

// what to return
func getUserByUsername(db *sql.DB, username string) *User {

	row := db.QueryRow(`SELECT "ID", username, password FROM "user" WHERE username = $1`, username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			// If no user with the given username is found, return nil
			log.Println("here username not found")
			log.Fatal(err)

		} else {
			log.Println("here username found")
		}
		// Handle other potential errors (e.g., database connection errors)
		// You may want to log the error or return an appropriate error response

	}
	log.Println(user.Username)
	log.Println(user.Password)
	return &user

}
