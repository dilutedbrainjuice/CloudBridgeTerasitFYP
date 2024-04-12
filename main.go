package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	db := ConnectPostgresDB()

	http.HandleFunc("/home/", homehandler)
	http.HandleFunc("/registersite/", registersitehandler)
	http.HandleFunc("/registerform/", RegisterHandler(db))
	http.HandleFunc("/about/", abouthandler)
	http.HandleFunc("/userdashboard/", validateToken(dashboardHandler))
	//Login page and login logic
	http.HandleFunc("/login/", loginhandler)
	http.HandleFunc("/loginform/", loginformhandler(db))
	http.HandleFunc("/logout/", logoutHandler)

	//initializing server
	log.Println("Listening on port 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))

}
