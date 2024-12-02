package model

import (
	"errors"
	"net/http"
)

type TreatmentRequest struct {
	VisitID    uint   `json:"visit_id"`
	Medicament string `json:"medicament"`
	Dose       string `json:"dose"`
}

func (t *TreatmentRequest) Bind(r *http.Request) error {
	if t.VisitID == 0 || t.Medicament == "" || t.Dose == "" {
		return errors.New("invalid input(s): tous les champs doivent être remplis")
	}
	return nil
}

type TreatmentResponse struct {
	ID         uint   `json:"id"`
	VisitID    uint   `json:"visit_id"`
	Medicament string `json:"medicament"`
	Dose       string `json:"dose"`
}
