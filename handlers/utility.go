package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeAsJSON(w http.ResponseWriter, o interface{}) {
	header := w.Header()
	header.Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(o)
}

func todo(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	fmt.Fprintf(w, "%s is not implemented", method)
}
