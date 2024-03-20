package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var filename = "index.html"
	t, err := template.ParseFiles(filename)

	if err != nil {
		fmt.Println("Error when parsing file", err)
		return
	}

	err = t.ExecuteTemplate(w, filename, nil)
	if err != nil {
		fmt.Println("Error when executing template", err)
		return
	}

}

// not working, figuring out a way to load dynamic content and change page with minimal pages
func loginHandler(w http.ResponseWriter, r *http.Request) { // not reading login.html
	wd, err := os.Getwd()

	if err != nil {
		fmt.Println("wd Error when executing template", err)
	}

	t, err := template.ParseFiles(wd + "/pages/login.html")

	if err != nil {
		fmt.Println("Error when parsing file", err)
		return
	}

	err = t.ExecuteTemplate(w, wd+"/pages/login.html", nil)

	if err != nil {
		fmt.Println("Error when executing template", err)
		return
	}

}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)

	case "/login":
		loginHandler(w, r)
	default:
		homeHandler(w, r)
	}
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server is listening on port 8080")
	server.ListenAndServe()

}
