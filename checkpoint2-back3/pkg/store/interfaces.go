package store

import "checkpoint2/internal/domain"

type DentistStoreInterface interface {
	ReadById(id int) (domain.Dentist, error)
	ReadAll() ([]domain.Dentist, error)
	ReadByRegistration(registration string) (domain.Dentist, error)
	Create(dentist domain.Dentist) (domain.Dentist, error)
	Update(id int, dentist domain.Dentist) (domain.Dentist, error)
	Patch(id int, dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type PatientStoreInterface interface {
	ReadById(id int) (domain.Patient, error)
	ReadAll() ([]domain.Patient, error)
	ReadByRg(rg string) (domain.Patient, error)
	Create(patient domain.Patient) (domain.Patient, error)
	Update(id int, patient domain.Patient) (domain.Patient, error)
	Patch(id int, patient domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type AppointmentStoreInterface interface {
	ReadById(id int) (domain.Appointment, error)
	ReadByRg(rg string) ([]domain.Appointment, error)
	CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error)
	CreateByRgAndRegistration(appointment domain.Appointment, rgPatient string, registrationDentist string) (domain.Appointment, error)
	Update(id int, appointment domain.Appointment) (domain.Appointment, error)
	Patch(id int, appointment domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}