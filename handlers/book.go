package handlers

import (
	"net/http"

	"runmarlowe.com/go-server-fun/services"
)

func getBooks(w http.ResponseWriter, r *http.Request) {
	s := &services.BookService{}
	book := s.Get()
	writeAsJSON(w, book)
}

// BookHandler returns Book data
func BookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getBooks(w, r)
	case http.MethodPut:
		todo(w, r)
	default:
		http.Error(w, "Unsupported method", http.StatusBadRequest)
	}
}
