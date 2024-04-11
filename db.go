package main

import (
	"database/sql"
	"fmt"
	"log"
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
