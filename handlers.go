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

	"golang.org/x/crypto/bcrypt"
)

func homehandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/base.html"))
	tmpl.Execute(w, nil)

}

func registersitehandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/index.html"))
	tmpl.Execute(w, nil)
}

func RegisterHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		// Extracting the file extension from the uploaded file
		fileExtension := filepath.Ext(handler.Filename)

		// Constructing the new filename using the username and file extension
		newFilename := newUser.Username + "_profilepic" + fileExtension

		// Creating the destination file with the new filename
		dst, err := os.Create("./uploads/" + newFilename)
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

		pk := insertUser(db, newUser)

		fmt.Printf("ID: %d\n", pk)

		http.Redirect(w, r, "/home/", http.StatusSeeOther)

	}

}

func abouthandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/about.html"))
	tmpl.Execute(w, nil)
}

func loginhandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/login.html"))
	tmpl.Execute(w, nil)
}

func LoginFormHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user User
		var verifyuser User
		user.Username = r.FormValue("username")
		log.Println(user.Username)
		user.Password = r.FormValue("password")
		log.Println(user.Password)

		// Get the existing entry present in the database for the given username
		result := db.QueryRow(`select password from "user" where username=$1`, user.Username)
		log.Println(result)

		err := result.Scan(verifyuser.Password)
		if err != nil {
			// If an entry with the username does not exist, send an "Unauthorized"(401) status
			if err == sql.ErrNoRows {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// If the error is of any other type, send a 500 status
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Compare the stored hashed password, with the hashed version of the password that was received
		if err = bcrypt.CompareHashAndPassword([]byte(verifyuser.Password), []byte(user.Password)); err != nil {
			// If the two passwords don't match, return a 401 status
			w.WriteHeader(http.StatusUnauthorized)
		}

	}
}
