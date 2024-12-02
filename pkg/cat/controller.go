package cat

import (
	"net/http"
	"strconv"
	"vet-clinic-api/config"
	"vet-clinic-api/database/dbmodels"
	model "vet-clinic-api/pkg/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// Enregistrer un nouveau profil de chat
func createCat(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cat model.CatRequest
		if err := render.Bind(r, &cat); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		newCat := &dbmodels.Cat{
			Nom:   cat.Nom,
			Age:   cat.Age,
			Race:  cat.Race,
			Poids: cat.Poids,
		}

		if err := cfg.CatRepository.Create(newCat); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, newCat)
	}
}

// Afficher tous les profils de chats
func getAllCats(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cats, err := cfg.CatRepository.FindAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, cats)
	}
}

// Afficher le profil d'un chat en fonction de son id
func getCatByID(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		cat, err := cfg.CatRepository.FindByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		render.JSON(w, r, cat)
	}
}

// Modifier le profil d'un chat
func updateCat(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		cat, err := cfg.CatRepository.FindByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		// Créer une instance de CatRequest pour le binding
		var catRequest model.CatRequest
		if err := render.Bind(r, &catRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Mettre à jour l'objet `cat` avec les nouvelles valeurs de `catRequest`
		cat.Nom = catRequest.Nom
		cat.Age = catRequest.Age
		cat.Race = catRequest.Race
		cat.Poids = catRequest.Poids
		// Mettre à jour l'objet dans la base de données
		if err := cfg.CatRepository.Update(cat); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, cat)
	}
}

// Supprimer le profil d'un chat avec son id
func deleteCat(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		if err := cfg.CatRepository.Delete(uint(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

// Afficher l'historique des consultations et des traitements d'un chat
func getCatHistory(cfg *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		catIDStr := chi.URLParam(r, "id")
		catID, err := strconv.Atoi(catIDStr)
		if err != nil {
			http.Error(w, "Invalid cat ID", http.StatusBadRequest)
			return
		}
		// Appeler la méthode pour obtenir l'historique complet des visites
		visits, err := cfg.CatRepository.FindHistoryByCatID(uint(catID))
		if err != nil {
			http.Error(w, "Failed to fetch cat history", http.StatusInternalServerError)
			return
		}
		// Retourner les visites dans la réponse
		render.JSON(w, r, map[string]interface{}{
			"visits": visits,
		})
	}
}
