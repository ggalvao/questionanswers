package util

import (
	"encoding/json"
	"net/http"
)

// JSONResponse marshals a type to JSON and writes it as response
func JSONResponse(w http.ResponseWriter, code int, output interface{}) {
	// Convert our interface to JSON
	response, _ := json.Marshal(output)
	// Set the content type to json for browsers
	w.Header().Set("Content-Type", "application/json")
	// Our response code
	w.WriteHeader(code)
	w.Write(response)
}
