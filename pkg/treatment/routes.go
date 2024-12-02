package treatment

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(cfg *config.Config) chi.Router {
	r := chi.NewRouter()
	r.Post("/createTreatment", createTreatment(cfg))
	r.Get("/getTreatments", getAllTreatments(cfg))
	r.Get("/getTreatment/{id}", getTreatmentByID(cfg))
	r.Put("/updateTreatment/{id}", updateTreatment(cfg))
	r.Delete("/deleteTreatment/{id}", deleteTreatment(cfg))
	r.Get("/history/getAllTreatmentsByVisit/{id}", getAllTreatmentsByVisit(cfg))
	return r
}
