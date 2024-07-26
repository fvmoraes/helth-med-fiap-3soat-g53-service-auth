package domain

type DoctorRepository interface {
	Create(doctor Doctor) error
	FindByEmail(email string) (*Doctor, error)
}

type PatientRepository interface {
	Create(patient Patient) error
	FindByEmail(email string) (*Patient, error)
}
