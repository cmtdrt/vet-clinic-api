package dbmodels

import (
	"gorm.io/gorm"
)

type Treatment struct {
	ID         uint   `gorm:"primaryKey"`
	VisitID    uint   `json:"visit_id"`
	Medicament string `json:"medicament"`
	Dose       string `json:"dose"`
}

type TreatmentRepository interface {
	Create(treatment *Treatment) error
	FindByID(id uint) (*Treatment, error)
	FindAll() ([]Treatment, error)
	Update(treatment *Treatment) error
	Delete(id uint) error
	FindAllByVisitID(visitID uint) ([]Treatment, error)
}

type treatmentRepository struct {
	db *gorm.DB
}

func NewTreatmentRepository(db *gorm.DB) TreatmentRepository {
	return &treatmentRepository{db}
}

func (r *treatmentRepository) Create(treatment *Treatment) error {
	return r.db.Create(treatment).Error
}

func (r *treatmentRepository) FindByID(id uint) (*Treatment, error) {
	var treatment Treatment
	if err := r.db.First(&treatment, id).Error; err != nil {
		return nil, err
	}
	return &treatment, nil
}

func (r *treatmentRepository) FindAll() ([]Treatment, error) {
	var treatments []Treatment
	if err := r.db.Find(&treatments).Error; err != nil {
		return nil, err
	}
	return treatments, nil
}

func (r *treatmentRepository) Update(treatment *Treatment) error {
	return r.db.Save(treatment).Error
}

func (r *treatmentRepository) Delete(id uint) error {
	return r.db.Delete(&Treatment{}, id).Error
}

func (r *treatmentRepository) FindAllByVisitID(visitID uint) ([]Treatment, error) {
	var treatments []Treatment
	if err := r.db.Where("visit_id = ?", visitID).Find(&treatments).Error; err != nil {
		return nil, err
	}
	return treatments, nil
}
