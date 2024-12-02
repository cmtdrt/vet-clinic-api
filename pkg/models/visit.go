package model

import (
	"errors"
	"net/http"
)

type VisitRequest struct {
	CatID       uint   `json:"cat_id"`
	Date        string `json:"date"`
	Motif       string `json:"motif"`
	Veterinaire string `json:"veterinaire"`
}

func (v *VisitRequest) Bind(r *http.Request) error {
	if v.Date == "" {
		return errors.New("date est requise")
	}
	if v.Motif == "" {
		return errors.New("motif est requis")
	}
	if v.Veterinaire == "" {
		return errors.New("veterinaire est requis")
	}
	return nil
}

type VisitResponse struct {
	ID          uint   `json:"id"`
	CatID       uint   `json:"cat_id"`
	Date        string `json:"date"`
	Motif       string `json:"motif"`
	Veterinaire string `json:"veterinaire"`
}
