package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

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

	//Handlers
	homehandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseGlob("templates/index.html"))
		tmpl.Execute(w, nil)
	}

	registerhandler := func(w http.ResponseWriter, r *http.Request) {
		log.Print("HTMX request received")

		// Get the form values
		name := r.PostFormValue("username")
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		//profilePic := r.FormFile("profilePic")
		city := r.PostFormValue("city")
		pcSpecs := r.PostFormValue("pcSpecs")
		description := r.PostFormValue("description")
		cloudService := r.PostFormValue("cloudService")

		// Log the user details
		log.Println("Username:", name)
		log.Println("Email:", email)
		log.Println("Password:", password)
		//log.Println("Profile Picture:", profilePic)
		log.Println("City:", city)
		log.Println("PC Specs:", pcSpecs)
		log.Println("Description:", description)
		log.Println("Cloud Service:", cloudService)

		http.Redirect(w, r, "/home/", http.StatusSeeOther)

	}

	http.HandleFunc("/home/", homehandler)
	http.HandleFunc("/registerform/", registerhandler)

	//initializing server
	log.Println("Listening on port 9000")

	log.Fatal(http.ListenAndServe(":9000", nil))

}
