package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("monkeygorilla")

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
		err := checkusernameHandler(db, r.PostFormValue("username"))
		if err != sql.ErrNoRows {
			//no proceeddo baby
			//FLAG :: Handle "username already taken error"
			log.Println("username taken")
			http.Error(w, "Error: Username already taken", http.StatusUnauthorized)

		} else {
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
				defaultImagePath := "/home/terasit/repos/CloudBridgeTerasitFYP/uploads/defaultprofilepic/default.png"
				defaultImage, err := os.Open(defaultImagePath)
				if err != nil {
					http.Error(w, "Failed to open default image", http.StatusInternalServerError)
					return
				}
				defer defaultImage.Close()

				// Create a new file in the uploads directory
				newFilename := newUser.Username + "_profilepic" + filepath.Ext(defaultImagePath)
				// You can change this filename if needed
				dst, err := os.Create("./uploads/" + newFilename)
				if err != nil {
					http.Error(w, "Failed to create file on server", http.StatusInternalServerError)
					return
				}
				defer dst.Close()

				// Copy the default image to the destination file
				if _, err := io.Copy(dst, defaultImage); err != nil {
					http.Error(w, "Failed to save file", http.StatusInternalServerError)
					return
				}

				// Set the ProfilePicURL with the new filename
				newUser.ProfilePicURL = "./uploads/" + newFilename

			} else {
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
			}

			latitudeStr := r.PostFormValue("latitude")
			longitudeStr := r.PostFormValue("longitude")

			latitude, err := strconv.ParseFloat(latitudeStr, 64)
			if err != nil {
				return
			}

			longitude, err := strconv.ParseFloat(longitudeStr, 64)
			if err != nil {
				return
			}

			newUser.Latitude = latitude
			newUser.Longitude = longitude

			newUser.PCSpecs = r.PostFormValue("pcSpecs")
			newUser.Description = r.PostFormValue("description")
			newUser.CloudService = r.PostFormValue("cloudService")
			currentTime := time.Now()
			currentTimeString := currentTime.Format("2006-01-02 15:04:05")
			newUser.CreatedAt = currentTimeString

			log.Println("Username:", newUser.Username)
			log.Println("IsProvider:", newUser.IsProvider)
			log.Println("Password:", newUser.Password)
			log.Println("Latitude:", newUser.Latitude)
			log.Println("Longitude:", newUser.Longitude)

			log.Println("ID:", newUser.ID)
			log.Println("Description:", newUser.Description)
			log.Println("PCSPEC:", newUser.PCSpecs)
			log.Println("ProfilePicURL:", newUser.ProfilePicURL)
			log.Println("Cloud Service:", newUser.CloudService)
			log.Println("Created At:", newUser.CreatedAt)

			pk := insertUser(db, newUser)

			fmt.Printf("ID: %d\n", pk)

			http.Redirect(w, r, "/home/", http.StatusSeeOther)

		}

	}

}

func checkusernameHandler(db *sql.DB, username string) error {

	var verifyR User
	//check username
	row := db.QueryRow(`SELECT username FROM "user" WHERE username = $1`, username)
	err := row.Scan(&verifyR.Username)
	return err
	//if row.username coincides with checkusername == return username has been taken, please choose another one

}

func abouthandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/about.html"))
	tmpl.Execute(w, nil)
}

func loginhandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseGlob("templates/login.html"))
	tmpl.Execute(w, nil)
}

func loginformhandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
				return
			}

			username := r.Form.Get("username")
			password := r.Form.Get("password")

			user := authenticateUser(db, username, password)
			if user != nil {
				// If the user is authenticated, create a JWT token
				expirationTime := time.Now().Add(24 * time.Hour) // Example: Token expires in 24 hours
				claims := &Claims{
					Username: username,
					StandardClaims: jwt.StandardClaims{
						ExpiresAt: expirationTime.Unix(),
					},
				}
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenString, err := token.SignedString(jwtKey)
				if err != nil {
					http.Error(w, "Error generating token", http.StatusInternalServerError)
					return
				}

				// Set the token as a cookie or send it in the response body
				http.SetCookie(w, &http.Cookie{
					Name:     "token",
					Value:    tokenString,
					Expires:  expirationTime,
					Path:     "/",                   // Set the cookie path
					HttpOnly: true,                  // HTTP only cookie
					Secure:   true,                  // Set to true if served over HTTPS
					SameSite: http.SameSiteNoneMode, // Consider setting SameSite attribute appropriately
				})

				// Redirect to protected page
				log.Println("Logged in handler version")
				http.Redirect(w, r, "/about/", http.StatusFound)
				return
			} else {
				// Handle invalid credentials
				//FLAG :: Handle "Invalid credentials"
				http.Error(w, "Invalid username or password", http.StatusUnauthorized)
				return
			}

		}

	}

}

func dashboardHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		userData, err := getuserlocation(db)
		if err != nil {
			log.Fatal(err)
			http.Error(w, "Fucking hell cant get a pussy", http.StatusInternalServerError)
		}

		jsonData, err := json.Marshal(userData)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)
			return
		}

		// Set response content type to JSON
		w.Header().Set("Content-Type", "application/json")

		// Write JSON data to response writer
		if _, err := w.Write(jsonData); err != nil {
			http.Error(w, fmt.Sprintf("Error writing JSON: %v", err), http.StatusInternalServerError)
			return
		}

	}
}

func dashboardshow(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseGlob("templates/dashboard.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

// Example logout handler
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: "",
		// Set expiration to the past
		MaxAge: -1,  // Alternatively, you can set MaxAge to 0
		Path:   "/", // Set the cookie path
	})

	//FLAG :: Handle cookie expire message
	//you have been logged out please log in again
	log.Println("Logged out")

	http.Redirect(w, r, "/home/", http.StatusFound)
}
