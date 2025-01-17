package usecase

import (
	"helthmed-auth/internal/auth/domain"
	"helthmed-auth/internal/auth/interfaces/messaging"
)

type DoctorUseCase struct {
	repo domain.DoctorRepository
}

func NewDoctorUseCase(repo domain.DoctorRepository) *DoctorUseCase {
	return &DoctorUseCase{repo: repo}
}

func (u *DoctorUseCase) Register(doctor *domain.Doctor) error {
	return u.repo.Create(doctor)

	// Publish the event
	messaging.PublishDoctorRegistered(doctor)
	return nil
}
