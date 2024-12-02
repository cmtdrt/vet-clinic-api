package cat

import (
	"vet-clinic-api/config"

	"github.com/go-chi/chi/v5"
)

func Routes(cfg *config.Config) chi.Router {
	r := chi.NewRouter()
	r.Post("/createCat", createCat(cfg))
	r.Get("/getCats", getAllCats(cfg))
	r.Get("/getCat/{id}", getCatByID(cfg))
	r.Put("/updateCat/{id}", updateCat(cfg))
	r.Delete("/deleteCat/{id}", deleteCat(cfg))
	r.Get("/{id}/history", getCatHistory(cfg))
	return r
}
