package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func loginHandler() {

}

func registerHandler() {
}

func main() {

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    ":9000",
		Handler: mux,
	}
	fmt.Println("listening on port 9000")
	srv.ListenAndServe()
}
