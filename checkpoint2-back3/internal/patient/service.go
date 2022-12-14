package patient

import (
	"checkpoint2/internal/domain"
	"errors"
)

type Service interface {
	ReadById(id int) (domain.Patient, error)
	ReadAll() ([]domain.Patient, error)
	ReadByRg(rg string) (domain.Patient, error)
	Create(patient domain.Patient) (domain.Patient, error)
	Update(id int, patient domain.Patient) (domain.Patient, error)
	Patch(id int, patient domain.Patient) (domain.Patient, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) ReadById(id int) (domain.Patient, error) {
	patient, err := s.r.ReadById(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (s *service) ReadByRg(rg string) (domain.Patient, error) {
	patient, err := s.r.ReadByRg(rg)
	if err != nil {
		return domain.Patient{}, err
	}
	return patient, nil
}

func (s *service) ReadAll() ([]domain.Patient, error) {
	patients, err := s.r.ReadAll()
	if err != nil {
		return []domain.Patient{}, err
	}
	return patients, nil
}

func (s *service) Create(patient domain.Patient) (domain.Patient, error) {
	patients, err := s.ReadAll()
	if err != nil {
		return domain.Patient{}, err
	}

	for i := range patients {
		if patients[i].RG == patient.RG {
			return domain.Patient{}, errors.New("rg already exists")
		}
	}

	createdPatient, err := s.r.Create(patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return createdPatient, nil
}

func (s *service) Update(id int, patient domain.Patient) (domain.Patient, error) {
	patients, err := s.ReadAll()
	if err != nil {
		return domain.Patient{}, err
	}

	for i := range patients {
		if patients[i].RG == patient.RG {
			return domain.Patient{}, errors.New("rg already exists")
		}
	}

	updatedPatient, err := s.r.Update(id, patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return updatedPatient, nil
}

func (s *service) Patch(id int, patient domain.Patient) (domain.Patient, error) {
	patients, err := s.ReadAll()
	if err != nil {
		return domain.Patient{}, err
	}

	for i := range patients {
		if patients[i].RG == patient.RG {
			return domain.Patient{}, errors.New("rg already exists")
		}
	}

	updatedPatient, err := s.r.Patch(id, patient)
	if err != nil {
		return domain.Patient{}, err
	}
	return updatedPatient, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}