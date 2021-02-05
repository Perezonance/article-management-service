package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/articles", articleHandler)
	r.HandleFunc("/articles/{articleID}", articleByIDHandler)
	r.HandleFunc("/articles/{userID}", articleByUserIDHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
