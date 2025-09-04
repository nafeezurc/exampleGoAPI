package main

// using logrus to log errors for debugging
import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/nafeezurc/exampleGoAPI/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	// creating error logger
	log.SetReportCaller(true)
	// creating a mux type variable, kind of like a struct
	var r *chi.Mux = chi.NewRouter
	// set up router i.e. endpoint definitions we want
	handlers.Handler(r)

	fmt.Println("Starting GO API Service...")

	fmt.Println("SUCCESS!")

	// start server
	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}

}
