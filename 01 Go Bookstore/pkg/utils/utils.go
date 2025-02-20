// utils package provides utility functions for the application.
package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseBody parses JSON data from an HTTP request body into a specified struct.
// This function reads the request body, unmarshals the JSON data, and populates the provided interface.
// @param r *http.Request - the incoming HTTP request containing the JSON body.
// @param x interface{} - a pointer to the struct where the JSON data will be unmarshaled.
func ParseBody(r *http.Request, x interface{}) {
	// Read the entire body from the HTTP request
	if body, err := io.ReadAll(r.Body); err == nil {
		// Unmarshal the JSON data into the provided interface
		if err := json.Unmarshal(body, x); err != nil {
			// If unmarshaling fails, return (could also log the error or handle it differently)
			return
		}
	}
}
