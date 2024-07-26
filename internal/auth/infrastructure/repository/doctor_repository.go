package repository

import (
	"helthmed-auth/internal/auth/domain"
	"helthmed-auth/internal/auth/infrastructure/db"
)

type DoctorRepository struct{}

func NewDoctorRepository() domain.DoctorRepository {
	return &DoctorRepository{}
}

func (r *DoctorRepository) Create(doctor *domain.Doctor) error {
	return db.DB.Create(doctor).Error
}

func (r *DoctorRepository) FindByEmail(email string) (*domain.Doctor, error) {
	var doctor domain.Doctor
	result := db.DB.Where("email = ?", email).First(&doctor)
	return &doctor, result.Error
}
