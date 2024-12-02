package database

import (
	"log"
	"vet-clinic-api/database/dbmodels"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDatabase initialise la connexion à la base de données et effectue la migration des modèles.
func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("vet_clinic_api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	Migrate(DB)
	log.Println("Database connected and migrated")
}

// Migrate effectue la migration des modèles.
func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&dbmodels.Cat{}, &dbmodels.Visit{}, &dbmodels.Treatment{})
	if err != nil {
		log.Println("Error during migration:", err)
	} else {
		log.Println("Database migrated successfully")
	}
}
