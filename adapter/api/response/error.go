// Package response provides the success and error responses to callers for the entire application
package response

import (
	"encoding/json"
	"net/http"
)

// Error defines the structure of errors for http responses
type Error struct {
	statusCode int
	Errors     []string `json:"errors"`
}

// NewError creates new Error type
func NewError(err error, status int) *Error {
	return &Error{
		statusCode: status,
		Errors:     []string{err.Error()},
	}
}

// Send returns a response with JSON format
func (e Error) Send(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.statusCode)
	return json.NewEncoder(w).Encode(e)
}
