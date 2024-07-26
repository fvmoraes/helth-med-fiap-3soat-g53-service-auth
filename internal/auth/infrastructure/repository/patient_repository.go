package repository

import (
	"helthmed-auth/internal/auth/domain"
	"helthmed-auth/internal/auth/infrastructure/db"
)

type PatientRepository struct{}

func NewPatientRepository() domain.PatientRepository {
	return &PatientRepository{}
}

func (r PatientRepository) Create(patient domain.Patient) error {
	return db.DB.Create(patient).Error
}

func (r PatientRepository) FindByEmail(email string) (domain.Patient, error) {
	var patient domain.Patient
	result := db.DB.Where("email = ?", email).First(&patient)
	return &patient, result.Error
}
