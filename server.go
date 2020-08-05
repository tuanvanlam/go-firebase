package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handle(response http.ResponseWriter, _ *http.Request) {
	_, _ = fmt.Fprintln(response, "Up and running...")
}

func main() {
	router := mux.NewRouter()
	const port string = ":8080"

	router.HandleFunc("/", handle)
	router.HandleFunc("/posts", getPosts).Methods(http.MethodGet)
	router.HandleFunc("/posts", addPost).Methods(http.MethodPost)

	log.Println("Server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
