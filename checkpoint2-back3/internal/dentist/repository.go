package dentist

import (
	"checkpoint2/internal/domain"
	"checkpoint2/pkg/store"
)

type Repository interface {
	ReadById(id int) (domain.Dentist, error)
	ReadAll() ([]domain.Dentist, error)
	ReadByRegistration(registration string) (domain.Dentist, error)
	Create(dentist domain.Dentist) (domain.Dentist, error)
	Update(id int, dentist domain.Dentist) (domain.Dentist, error)
	Patch(id int, dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type repository struct {
	storage store.DentistStoreInterface
}

func NewRepository(storage store.DentistStoreInterface) Repository {
	return &repository{storage}
}

func (r *repository) ReadById(id int) (domain.Dentist, error) {
	dentist, err := r.storage.ReadById(id)
	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil
}

func (r *repository) ReadAll() ([]domain.Dentist, error) {
	dentists, err := r.storage.ReadAll()
	if err != nil {
		return []domain.Dentist{}, err
	}
	return dentists, nil
}

func (r *repository) ReadByRegistration(registration string) (domain.Dentist, error) {
	dentist, err := r.storage.ReadByRegistration(registration)
	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil
}

func (r *repository) Create(d domain.Dentist) (domain.Dentist, error) {
	dentist, err := r.storage.Create(d)
	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil
}

func (r *repository) Update(id int, d domain.Dentist) (domain.Dentist, error) {
	dentist, err := r.storage.Update(id, d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (r *repository) Patch(id int, d domain.Dentist) (domain.Dentist, error) {
	dentist, err := r.storage.Patch(id, d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return dentist, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}