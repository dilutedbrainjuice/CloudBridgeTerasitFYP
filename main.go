package main

import (
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	db := ConnectPostgresDB()
	//Handlers
	homehandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseGlob("templates/index.html"))
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/home/", homehandler)
	http.HandleFunc("/registerform/", RegisterHandler(db))

	//initializing server
	log.Println("Listening on port 9000")

	log.Fatal(http.ListenAndServe(":9000", nil))

}
