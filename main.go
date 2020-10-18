package main

import (
	"log"
	"net/http"

	"runmarlowe.com/go-server-fun/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", handlers.BookHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
