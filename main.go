package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func loginHandler() {

}

func registerHandler() {
}

type User struct {
}

func main() {

	//database
	var db *sql.DB
	var err error
	db, err = sql.Open("postgres", "user=postgres password=admin@123 dbname=cloudbridge sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database:", err)
		return
	}
	fmt.Println("Successfully connected to cloudbridge database")

	//initializing server
	mux := http.NewServeMux()
	srv := &http.Server{
		Addr:    ":9000",
		Handler: mux,
	}
	fmt.Println("listening on port 9000")
	srv.ListenAndServe()
}
