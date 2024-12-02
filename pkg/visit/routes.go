package visit

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(cfg *config.Config) chi.Router {
	r := chi.NewRouter()
	r.Post("/createVisit", createVisit(cfg))
	r.Get("/getVisits", getAllVisits(cfg))
	r.Get("/getVisit/{id}", getVisitByID(cfg))
	r.Put("/updateVisit/{id}", updateVisit(cfg))
	r.Delete("/deleteVisit/{id}", deleteVisit(cfg))
	r.Get("/history/getAllVisitsByCat/{id}", getAllVisitsByCat(cfg))
	r.Get("/withDate/{date}", getVisitsByDate(cfg))
	r.Get("/withMotif/{motif}", getVisitsByMotif(cfg))
	r.Get("/withVeterinaire/{veterinaire}", getVisitsByVeterinaire(cfg))
	return r
}
