package config

import (
	"vet-clinic-api/database"
	"vet-clinic-api/database/dbmodels"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	CatRepository       dbmodels.CatRepository
	VisitRepository     dbmodels.VisitRepository
	TreatmentRepository dbmodels.TreatmentRepository
}

func New() (*Config, error) {
	config := Config{}

	// Initialisation de la connexion à la base de données
	databaseSession, err := gorm.Open(sqlite.Open("vet_clinic_api.db"), &gorm.Config{})
	if err != nil {
		return &config, err
	}

	// Migration des modèles en utilisant le package database
	database.Migrate(databaseSession)

	// Initialisation des repositories
	config.CatRepository = dbmodels.NewCatRepository(databaseSession)
	config.VisitRepository = dbmodels.NewVisitRepository(databaseSession)
	config.TreatmentRepository = dbmodels.NewTreatmentRepository(databaseSession)

	return &config, nil
}
