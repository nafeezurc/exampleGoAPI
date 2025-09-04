package middleware

import (
	"errors"
	"net/http"

	"github.com/nafeezurc/exampleGoAPI/api"
	"github.com/nafeezurc/exampleGoAPI/internal/tools"
	log "github.com/sirupsen/logrus"
)

// custom unauthorized error
var UnAuthorizedError = errors.New("Invalid username or token")

// needs to take in and return an http.Handler interface
func Authorization(next http.Handler) http.Handler {
	// return using anonymous function
	// HandlerFunc in http package, which takes in a ResponseWriter and a pointer to the request
	// ResponseWriter is what you use to send a response to the caller
	// Request contains all the information about the incoming request
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// grab username from request pointer
		var username string = r.URL.Query().Get("username")
		// get the authorization token
		var token = r.Header.Get("Authorization")
		var err error

		// if username or token is empty, return an error
		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}
	})

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}
}
