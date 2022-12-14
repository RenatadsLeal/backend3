package dentist

import (
	"checkpoint2/internal/domain"
	"errors"
)

type Service interface {
	ReadById(id int) (domain.Dentist, error)
	ReadAll() ([]domain.Dentist, error)
	ReadByRegistration(registration string) (domain.Dentist, error)
	Create(dentist domain.Dentist) (domain.Dentist, error)
	Update(id int, dentist domain.Dentist) (domain.Dentist, error)
	Patch(id int, dentist domain.Dentist) (domain.Dentist, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) ReadById(id int) (domain.Dentist, error) {
	dentist, err := s.r.ReadById(id)
	if err != nil {
		return domain.Dentist{}, err
	}

	return dentist, nil
}

func (s *service) ReadAll() ([]domain.Dentist, error) {
	dentists, err := s.r.ReadAll()
	if err != nil {
		return []domain.Dentist{}, err
	}
	return dentists, nil
}

func (s *service) ReadByRegistration(registration string) (domain.Dentist, error) {
	dentist, err := s.r.ReadByRegistration(registration)
	if err != nil {
		return domain.Dentist{}, err
	}
	
	return dentist, nil
}

func (s *service) Create(d domain.Dentist) (domain.Dentist, error) {
	dentists, err := s.ReadAll()
	if err != nil {
		return domain.Dentist{}, err
	}

	for i := range dentists {
		if dentists[i].Registration == d.Registration {
			return domain.Dentist{}, errors.New("registration already exists")
		}
	}

	newDentist, err := s.r.Create(d)
	if err != nil {
		return domain.Dentist{}, err
	}

	return newDentist, nil
}

func (s *service) Update(id int, d domain.Dentist) (domain.Dentist, error) {
	dentists, err := s.ReadAll()
	if err != nil {
		return domain.Dentist{}, err
	}

	for i := range dentists {
		if dentists[i].Registration == d.Registration {
			return domain.Dentist{}, errors.New("registration already exists")
		}
	}

	updatedDentist, err := s.r.Update(id, d)
	if err != nil {
		return domain.Dentist{}, err
	}
	
	return updatedDentist, nil
}

func (s *service) Patch(id int, dentist domain.Dentist) (domain.Dentist, error) {
	dentists, err := s.ReadAll()
	if err != nil {
		return domain.Dentist{}, err
	}

	for i := range dentists {
		if dentists[i].Registration == dentist.Registration {
			return domain.Dentist{}, errors.New("registration already exists")
		}
	}

	updatedDentist, err := s.r.Patch(id, dentist)
	if err != nil {
		return domain.Dentist{}, err
	}
	return updatedDentist, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}

	return nil
}