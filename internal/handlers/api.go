package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/nafeezurc/goapi/internal/middleware"
)

// middleware is a function that gets called before the primary function that handles the endpoint
func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes)

	// set up route
	// create an anonymous function middleware
	r.Route("/account", func(router chi.Router) {
		// middleware for /account route
		router.use(middleware.Authorization)
		// created an endpoint at /account/coins
		// in order to use GetCoinsBalance in the above Get method
		// we define it such that it takes in a response writer and a pointer to the request as parameters
		router.Get("/coins", GetCoinBalance)
	})
}
