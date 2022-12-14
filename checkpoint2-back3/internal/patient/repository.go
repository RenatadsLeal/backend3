package patient

import (
	"checkpoint2/internal/domain"
	"checkpoint2/pkg/store"
)

type Repository interface {
	ReadById(id int) (domain.Patient, error)
	ReadAll() ([]domain.Patient, error)
	ReadByRg(rg string) (domain.Patient, error)
	Create(patient domain.Patient) (domain.Patient, error)
	Update(id int, patient domain.Patient) (domain.Patient, error)
	Patch(id int, patient domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type repository struct {
	storage store.PatientStoreInterface
}

func NewRepository(storage store.PatientStoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) ReadById(id int) (domain.Patient, error) {
	patient, err := r.storage.ReadById(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (r *repository) ReadAll() ([]domain.Patient, error) {
	patients, err := r.storage.ReadAll()
	if err != nil {
		return []domain.Patient{}, err
	}
	return patients, nil
}

func (r *repository) ReadByRg(rg string) (domain.Patient, error) {
	patient, err := r.storage.ReadByRg(rg)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (r *repository) Create(p domain.Patient) (domain.Patient, error) {
	patient, err := r.storage.Create(p)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (r *repository) Update(id int, p domain.Patient) (domain.Patient, error) {
	patient, err := r.storage.Update(id, p)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (r *repository) Patch(id int, p domain.Patient) (domain.Patient, error) {
	patient, err := r.storage.Patch(id, p)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}