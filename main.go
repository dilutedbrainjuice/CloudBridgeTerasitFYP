package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID            int64  `json:"id,omitempty"`
	Username      string `json:"username"`
	IsProvider    bool   `json:"isProvider"`
	Email         string `json:"email"`
	Password      string `json:"password,omitempty"` // Hashed password, omit from JSON
	ProfilePicURL string `json:"profilePicURL"`
	City          string `json:"city"`
	PCSpecs       string `json:"pcSpecs"`
	Description   string `json:"description"`
	CloudService  string `json:"cloudService"`
	CreatedAt     string `json:"createdAt,omitempty"`
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

		var newUser User

		// Get the form values
		newUser.Username = r.PostFormValue("username")

		isProviderStr := r.PostFormValue("isprovider")
		var isProviderbool bool
		if isProviderStr == "true" {
			isProviderbool = true
		} else {
			isProviderbool = false
		}
		newUser.IsProvider = isProviderbool

		newUser.Email = r.PostFormValue("email")

		priorhash := r.PostFormValue("password")
		hashedpassword, err := bcrypt.GenerateFromPassword([]byte(priorhash), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		hashedPasswordString := string(hashedpassword)
		newUser.Password = hashedPasswordString

		profilePic, handler, err := r.FormFile("profilePic")
		if err != nil {
			http.Error(w, "Failed to retrieve file", http.StatusBadRequest)
			return
		}
		defer profilePic.Close()

		dst, err := os.Create("./uploads/" + handler.Filename)
		if err != nil {
			http.Error(w, "Failed to create file on server", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		// Copy the uploaded file to the destination file
		if _, err := io.Copy(dst, profilePic); err != nil {
			http.Error(w, "Failed to save file", http.StatusInternalServerError)
			return
		}

		// Get the file path
		// Extracting the file extension from the uploaded file
		fileExtension := filepath.Ext(handler.Filename)

		// Constructing the new filename using the username and file extension
		newFilename := newUser.Username + "_profilepic" + fileExtension

		// Setting the ProfilePicURL with the new filename
		newUser.ProfilePicURL = "./uploads/" + newFilename

		newUser.City = r.PostFormValue("city")
		newUser.PCSpecs = r.PostFormValue("pcSpecs")
		newUser.Description = r.PostFormValue("description")
		newUser.CloudService = r.PostFormValue("cloudService")
		currentTime := time.Now()
		currentTimeString := currentTime.Format("2006-01-02 15:04:05")
		newUser.CreatedAt = currentTimeString

		log.Println("Username:", newUser.Username)
		log.Println("IsProvider:", newUser.IsProvider)
		log.Println("Password:", newUser.Password)
		log.Println("ID:", newUser.ID)
		log.Println("Description:", newUser.Description)
		log.Println("PCSPEC:", newUser.PCSpecs)
		log.Println("ProfilePicURL:", newUser.ProfilePicURL)
		log.Println("City", newUser.City)
		log.Println("Cloud Service:", newUser.CloudService)
		log.Println("Created At:", newUser.CreatedAt)

		// Insert data into the database
		_, err = db.Query("INSERT INTO user(Username, IsProvider, Email, Password, ProfilePicUrl, City, PC_specs, Description, Cloud_service, CreatedAt) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
			newUser.Username, newUser.IsProvider, newUser.Email, newUser.Password, newUser.ProfilePicURL, newUser.City, newUser.PCSpecs, newUser.Description, newUser.CloudService, newUser.CreatedAt)
		if err != nil {
			log.Println("here")
			http.Error(w, "Failed to insert data into database", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/home/", http.StatusSeeOther)
		// insert insert to database function here

	}

	http.HandleFunc("/home/", homehandler)
	http.HandleFunc("/registerform/", registerhandler)

	//initializing server
	log.Println("Listening on port 9000")

	log.Fatal(http.ListenAndServe(":9000", nil))

}
