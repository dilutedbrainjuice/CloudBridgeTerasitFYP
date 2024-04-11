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
	fmt.Printf("\n%s\n", user.Username)
	fmt.Printf("%s\n", user.Password)
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

	row := db.QueryRow(`SELECT "ID", "Username", "password" FROM "user" WHERE "Username" = ?`, username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			// If no user with the given username is found, return nil
			log.Println("here username not found")
			log.Fatal(err)
		}
		// Handle other potential errors (e.g., database connection errors)
		// You may want to log the error or return an appropriate error response
		log.Println("here username found")
	}

	return &user

}
