package main

import (
	"github.com/golang/fosterTransport/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func init() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Home)

	http.Handle("/", router)

	log.Println("Listening...")
	http.ListenAndServe(":80", nil)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.Home)

	http.Handle("/", router)

	log.Println("Listening...")
	http.ListenAndServe(":80", nil)
}
