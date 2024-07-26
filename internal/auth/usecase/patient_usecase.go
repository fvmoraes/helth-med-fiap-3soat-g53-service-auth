package usecase

import (
	"helthmed-auth/internal/auth/domain"
	"helthmed-auth/internal/auth/interfaces/messaging"
)

type PatientUseCase struct {
	repo domain.PatientRepository
}

func NewPatientUseCase(repo domain.PatientRepository) *PatientUseCase {
	return &PatientUseCase{repo: repo}
}

func (u *PatientUseCase) Register(patient *domain.Patient) error {
	err := u.repo.Create(patient)
	// if err != nil {
	// 	return err
	// }

	// Publish the event
	messaging.PublishPatientRegistered(patient)
	return nil
}
