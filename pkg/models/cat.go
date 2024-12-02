package model

import (
	"errors"
	"net/http"
)

type CatRequest struct {
	Nom   string  `json:"nom"`
	Age   int     `json:"age"`
	Race  string  `json:"race"`
	Poids float64 `json:"poids"`
}

func (c *CatRequest) Bind(r *http.Request) error {
	if c.Nom == "" {
		return errors.New("nom est requis")
	}
	if c.Race == "" {
		return errors.New("la race est requis")
	}
	if c.Age < 0 {
		return errors.New("age doit être un entier positif")
	}
	if c.Poids < 0 {
		return errors.New("poids doit être une valeur positive")
	}
	return nil
}

type CatResponse struct {
	ID    uint    `json:"id"`
	Nom   string  `json:"nom"`
	Age   int     `json:"age"`
	Race  string  `json:"race"`
	Poids float64 `json:"poids"`
}
