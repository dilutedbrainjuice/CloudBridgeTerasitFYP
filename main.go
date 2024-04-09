package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func loginHandler() {

}

func registerHandler() {
}

type User struct {
	ID            int64     `json:"id,omitempty"`
	Username      string    `json:"username"`
	IsProvider    bool      `json:"isProvider"`
	Email         string    `json:"email"`
	Password      string    `json:"password,omitempty"` // Hashed password, omit from JSON
	ProfilePicURL string    `json:"profilePicURL"`
	City          string    `json:"city"`
	PCSpecs       string    `json:"pcSpecs"`
	Description   string    `json:"description"`
	CloudService  string    `json:"cloudService"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
}

type Message struct {
	MessageID  int64     `json:"messageID,omitempty"`
	SenderID   int64     `json:"senderID"`
	ReceiverID int64     `json:"receiverID"`
	Message    string    `json:"message"`
	SentAt     time.Time `json:"sentAt,omitempty"`
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
