package treatment

import (
	"net/http"
	"strconv"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodels"
	model "vet-clinic-api/pkg/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Créer un nouveau traitement
func createTreatment(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var treatment model.TreatmentRequest
		if err := render.Bind(r, &treatment); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newTreatment := &dbmodels.Treatment{
			VisitID:    treatment.VisitID,
			Medicament: treatment.Medicament,
			Dose:       treatment.Dose,
		}

		if err := cfg.TreatmentRepository.Create(newTreatment); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, newTreatment)
	}
}

// Afficher tous les traitements
func getAllTreatments(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		treatments, err := cfg.TreatmentRepository.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, treatments)
	}
}

// Afficher un traitement en fonction de son id
func getTreatmentByID(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		treatment, err := cfg.TreatmentRepository.FindByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		render.JSON(w, r, treatment)
	}
}

// Modifier un traitement en fonction de son id
func updateTreatment(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		treatment, err := cfg.TreatmentRepository.FindByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		// Créer une instance de TreatmentRequest pour le binding
		var treatmentRequest model.TreatmentRequest
		if err := render.Bind(r, &treatmentRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Mettre à jour l'objet `treatment` avec les nouvelles valeurs de `treatmentRequest`
		treatment.VisitID = treatmentRequest.VisitID
		treatment.Medicament = treatmentRequest.Medicament
		treatment.Dose = treatmentRequest.Dose
		// Mettre à jour l'objet dans la base de données
		if err := cfg.TreatmentRepository.Update(treatment); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, treatment)
	}
}

// Supprimer un traitement en fonction de son id
func deleteTreatment(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		if err := cfg.TreatmentRepository.Delete(uint(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

// Afficher tous les traitements d'une visite
func getAllTreatmentsByVisit(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			http.Error(w, "Invalid visit ID", http.StatusBadRequest)
			return
		}

		treatments, err := cfg.TreatmentRepository.FindAllByVisitID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, treatments)
	}
}
