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
	query := `INSERT INTO "user" (username, isprovider, password, profilepicurl, pc_specs, description, cloud_service, createdat, latitude, longitude) VALUES ( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING "ID"`

	var pk int
	err := db.QueryRow(query, user.Username, user.IsProvider, user.Password, user.ProfilePicURL, user.PCSpecs, user.Description, user.CloudService, user.CreatedAt, user.Latitude, user.Longitude).Scan(&pk)
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

func getuserlocation(db *sql.DB) ([]User, error) {

	rows, err := db.Query(`SELECT "ID", username, isprovider, latitude, longitude FROM "user"`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Username, &user.IsProvider, &user.Latitude, &user.Longitude)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil

}
