package dbmodels

import (
	"gorm.io/gorm"
)

type Cat struct {
	ID    uint    `gorm:"primaryKey"`
	Nom   string  `json:"nom"`
	Age   int     `json:"age"`
	Race  string  `json:"race"`
	Poids float64 `json:"poids"`
}

type CatRepository interface {
	Create(cat *Cat) error
	FindByID(id uint) (*Cat, error)
	FindAll() ([]Cat, error)
	Update(cat *Cat) error
	Delete(id uint) error
	FindHistoryByCatID(catID uint) ([]Visit, error)
}

type catRepository struct {
	db *gorm.DB
}

func NewCatRepository(db *gorm.DB) CatRepository {
	return &catRepository{db}
}

// Créer un nouveau profil de chat
func (r *catRepository) Create(cat *Cat) error {
	return r.db.Create(cat).Error
}

// Trouver le profil d'un chat avec son id
func (r *catRepository) FindByID(id uint) (*Cat, error) {
	var cat Cat
	if err := r.db.First(&cat, id).Error; err != nil {
		return nil, err
	}
	return &cat, nil
}

// Afficher tous les profils de chat existants
func (r *catRepository) FindAll() ([]Cat, error) {
	var cats []Cat
	if err := r.db.Find(&cats).Error; err != nil {
		return nil, err
	}
	return cats, nil
}

// Modifier le profil d'un chat
func (r *catRepository) Update(cat *Cat) error {
	return r.db.Save(cat).Error
}

// Supprimer le profil d'un chat
func (r *catRepository) Delete(id uint) error {
	return r.db.Delete(&Cat{}, id).Error
}

// Afficher l'historique des consultations et traitements d'un chat
func (r *catRepository) FindHistoryByCatID(catID uint) ([]Visit, error) {
	var visits []Visit

	// Récupérer toutes les visites pour le chat, en préchargeant les traitements associés
	if err := r.db.Preload("Traitements").Where("cat_id = ?", catID).Find(&visits).Error; err != nil {
		return nil, err
	}

	return visits, nil
}
