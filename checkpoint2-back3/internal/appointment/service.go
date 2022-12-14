package appointment

import (
	"checkpoint2/internal/domain"
)

type Service interface {
	ReadById(id int) (domain.Appointment, error)
	ReadByRg(rg string) ([]domain.Appointment, error)
	CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error)
	CreateByRgAndRegistration(appointment domain.Appointment, rgPatient string, registrationDentist string) (domain.Appointment, error)
	Update(id int, appointment domain.Appointment) (domain.Appointment, error)
	Patch(id int, appointment domain.Appointment) (domain.Appointment, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) ReadById(id int) (domain.Appointment, error) {
	appointment, err := s.r.ReadById(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *service) ReadByRg(rg string) ([]domain.Appointment, error) {
	appointments, err := s.r.ReadByRg(rg)
	if err != nil {
		return []domain.Appointment{}, err
	}
	return appointments, nil
}

func (s *service) CreateById(a domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error) {
	appointment, err := s.r.CreateById(a, idPatient, idDentist)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *service) CreateByRgAndRegistration(a domain.Appointment, rgPatient string, registrationDentist string) (domain.Appointment, error) {
	appointment, err := s.r.CreateByRgAndRegistration(a, rgPatient, registrationDentist)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *service) Update(id int, a domain.Appointment) (domain.Appointment, error) {
	appointment, err := s.r.Update(id, a)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *service) Patch(id int, a domain.Appointment) (domain.Appointment, error) {
	appointment, err := s.r.Patch(id, a)
	if err != nil {
		return domain.Appointment{}, err
	}
	return appointment, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}

	return nil
}