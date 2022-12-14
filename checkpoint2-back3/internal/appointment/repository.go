package appointment

import (
	"checkpoint2/internal/domain"
	"checkpoint2/pkg/store"
)

type Repository interface {
	ReadById(id int) (domain.Appointment, error)
	ReadByRg(rg string) ([]domain.Appointment, error)
	CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error)
	CreateByRgAndRegistration(appointment domain.Appointment, rgPatient string, registrationDentist string) (domain.Appointment, error)
	Update(id int, appointment domain.Appointment) (domain.Appointment, error)
	Patch(id int, appointment domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}

type repository struct {
	storage store.AppointmentStoreInterface
}

func NewRepository(storage store.AppointmentStoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) ReadById(id int) (domain.Appointment, error) {
	appointment, err := r.storage.ReadById(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (r *repository) ReadByRg(rg string) ([]domain.Appointment, error) {
	appointments, err := r.storage.ReadByRg(rg)
	if err != nil {
		return []domain.Appointment{}, err
	}
	return appointments, nil
}

func (r *repository) CreateById(a domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error) {
	appointment, err := r.storage.CreateById(a, idPatient, idDentist)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (r *repository) CreateByRgAndRegistration(appointment domain.Appointment, rgPatient string, registrationDentist string) (domain.Appointment, error) {
	createdAppointment, err := r.storage.CreateByRgAndRegistration(appointment, rgPatient, registrationDentist)
	if err != nil {
		return domain.Appointment{}, err
	}
	return createdAppointment, nil
}

func (r *repository) Update(id int, a domain.Appointment) (domain.Appointment, error) {
	appointment, err := r.storage.Update(id, a)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (r *repository) Patch(id int, a domain.Appointment) (domain.Appointment, error) {
	appointment, err := r.storage.Patch(id, a)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}