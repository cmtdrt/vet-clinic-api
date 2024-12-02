package main

import (
	"log"
	"net/http"
	"vet-clinic-api/config"
	"vet-clinic-api/pkg/cat"
	"vet-clinic-api/pkg/treatment"
	"vet-clinic-api/pkg/visit"

	"github.com/go-chi/chi/v5"
)

func Routes(cfg *config.Config) *chi.Mux {
	r := chi.NewRouter()
	r.Mount("/api/cat", cat.Routes(cfg))
	r.Mount("/api/visit", visit.Routes(cfg))
	r.Mount("/api/treatment", treatment.Routes(cfg))
	return r
}

func main() {
	// Initialisation de la configuration
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error:", err)
	}

	// Initialisation des routes
	router := Routes(configuration)

	log.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
