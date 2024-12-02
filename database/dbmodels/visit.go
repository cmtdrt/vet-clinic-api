package dbmodels

import (
	"gorm.io/gorm"
)

type Visit struct {
	ID          uint        `gorm:"primaryKey"`
	CatID       uint        `json:"cat_id"`
	Date        string      `json:"date"`
	Motif       string      `json:"motif"`
	Veterinaire string      `json:"veterinaire"`
	Traitements []Treatment `gorm:"foreignKey:VisitID"` // Relation avec les traitements, optionnelle
}

type VisitRepository interface {
	Create(visit *Visit) error
	FindByID(id uint) (*Visit, error)
	FindAll() ([]Visit, error)
	Update(visit *Visit) error
	Delete(id uint) error
	FindAllByCatID(catID uint) ([]Visit, error)
	FindByDate(date string) ([]Visit, error)
	FindByMotif(motif string) ([]Visit, error)
	FindByVeterinaire(veterinaire string) ([]Visit, error)
}

type visitRepository struct {
	db *gorm.DB
}

func NewVisitRepository(db *gorm.DB) VisitRepository {
	return &visitRepository{db}
}

func (v *Visit) TableName() string {
	return "visits"
}

func (r *visitRepository) Create(visit *Visit) error {
	return r.db.Create(visit).Error
}

func (r *visitRepository) FindAll() ([]Visit, error) {
	var visits []Visit
	if err := r.db.Preload("Traitements").Find(&visits).Error; err != nil {
		return nil, err
	}
	return visits, nil
}

func (r *visitRepository) FindByID(id uint) (*Visit, error) {
	var visit Visit
	if err := r.db.Preload("Traitements").First(&visit, id).Error; err != nil {
		return nil, err
	}
	return &visit, nil
}

func (r *visitRepository) Update(visit *Visit) error {
	return r.db.Save(visit).Error
}

func (r *visitRepository) Delete(id uint) error {
	return r.db.Delete(&Visit{}, id).Error
}

func (r *visitRepository) FindAllByCatID(catID uint) ([]Visit, error) {
	var visits []Visit
	if err := r.db.Preload("Traitements").Where("cat_id = ?", catID).Find(&visits).Error; err != nil {
		return nil, err
	}
	return visits, nil
}

// Recherche par date avec les traitements associés
func (r *visitRepository) FindByDate(date string) ([]Visit, error) {
	var visits []Visit
	if err := r.db.Preload("Traitements").Where("date = ?", date).Find(&visits).Error; err != nil {
		return nil, err
	}
	return visits, nil
}

// Recherche par motif avec les traitements associés
func (r *visitRepository) FindByMotif(motif string) ([]Visit, error) {
	var visits []Visit
	if err := r.db.Preload("Traitements").Where("motif LIKE ?", "%"+motif+"%").Find(&visits).Error; err != nil {
		return nil, err
	}
	return visits, nil
}

// Recherche par vétérinaire avec les traitements associés
func (r *visitRepository) FindByVeterinaire(veterinaire string) ([]Visit, error) {
	var visits []Visit
	if err := r.db.Preload("Traitements").Where("veterinaire LIKE ?", "%"+veterinaire+"%").Find(&visits).Error; err != nil {
		return nil, err
	}
	return visits, nil
}
