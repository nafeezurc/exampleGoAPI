package api

import (
	"encoding/json"
	"net/http"
)

// Coin balance params
type CoinBalanceParams struct {
	Username string
}

// Coin balance response
type CoinBalanceResponse struct {
	// success code, usually 200
	Code int

	// account balance
	Balance int64
}

// Error response
type Error struct {
	// error code
	Code int

	// error message
	Message string
}

// function that will return error response to the person that called the endpoint
// takes in writer, error message, and error code
func writeError(w http.ResponseWriter, message string, code int) {
	// create an error struct
	resp := Error{
		Code:    code,
		Message: message,
	}

	// set content type, return a json for this case
	w.Header().Set("Content-Type", "application/json")
	// write the error code
	w.WriteHeader(code)

	// write the error struct out
	json.NewEncoder(w).Encode(resp)
}

// not calling above error directly in functions since we'll have a couple of different errors in our functions
// create a wrapper for error function
var (
	// take in response writer and error
	// use this to return a specific error in response
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	// use this for a generic error in response
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occured", http.StatusInternalServerError)
	}
)
