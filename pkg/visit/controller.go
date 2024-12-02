package visit

import (
	"net/http"
	"strconv"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodels"
	model "vet-clinic-api/pkg/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Créer une consultation
func createVisit(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var visit model.VisitRequest
		if err := render.Bind(r, &visit); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newVisit := &dbmodels.Visit{
			CatID:       visit.CatID,
			Date:        visit.Date,
			Motif:       visit.Motif,
			Veterinaire: visit.Veterinaire,
		}

		if err := cfg.VisitRepository.Create(newVisit); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, newVisit)
	}
}

// Récupérer toutes les consultations
func getAllVisits(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Charger les consultations avec les traitements associés
		visits, err := cfg.VisitRepository.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, visits)
	}
}

// Récupérer une consultation en fonction de son id
func getVisitByID(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		// Précharger les traitements associés à la consultation
		visit, err := cfg.VisitRepository.FindByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		render.JSON(w, r, visit)
	}
}

// Modifier une consultation
func updateVisit(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		visit, err := cfg.VisitRepository.FindByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		// Créer une instance de VisitRequest pour le binding
		var visitRequest model.VisitRequest
		if err := render.Bind(r, &visitRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Mettre à jour l'objet `visit` avec les nouvelles valeurs de `visitRequest`
		visit.CatID = visitRequest.CatID
		visit.Date = visitRequest.Date
		visit.Motif = visitRequest.Motif
		visit.Veterinaire = visitRequest.Veterinaire
		// Mettre à jour l'objet dans la base de données
		if err := cfg.VisitRepository.Update(visit); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, visit)
	}
}

// Supprimer une consultation
func deleteVisit(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		if err := cfg.VisitRepository.Delete(uint(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

// Recherche de l'historique d'un chat via son id
func getAllVisitsByCat(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid cat ID", http.StatusBadRequest)
			return
		}
		// Charger les consultations avec les traitements associés
		visits, err := cfg.VisitRepository.FindAllByCatID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, visits)
	}
}

// Recherche par critère
// Recherche des consultations en fonction de la date
func getVisitsByDate(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		date := chi.URLParam(r, "date")
		if date == "" {
			http.Error(w, "Date parameter is required", http.StatusBadRequest)
			return
		}
		// Charger les consultations avec les traitements associés
		visits, err := cfg.VisitRepository.FindByDate(date)
		if err != nil {
			http.Error(w, "Failed to fetch visits", http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, visits)
	}
}

// Recherche des consultations en fonction du motif
func getVisitsByMotif(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		motif := chi.URLParam(r, "motif")
		if motif == "" {
			http.Error(w, "Motif parameter is required", http.StatusBadRequest)
			return
		}
		// Charger les consultations avec les traitements associés
		visits, err := cfg.VisitRepository.FindByMotif(motif)
		if err != nil {
			http.Error(w, "Failed to fetch visits", http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, visits)
	}
}

// Recherche des consultations en fonction du vétérinaire
func getVisitsByVeterinaire(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		veterinaire := chi.URLParam(r, "veterinaire")
		if veterinaire == "" {
			http.Error(w, "Veterinaire parameter is required", http.StatusBadRequest)
			return
		}
		// Charger les consultations avec les traitements associés
		visits, err := cfg.VisitRepository.FindByVeterinaire(veterinaire)
		if err != nil {
			http.Error(w, "Failed to fetch visits", http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, visits)
	}
}
