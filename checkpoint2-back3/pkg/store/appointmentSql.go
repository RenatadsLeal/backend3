package store

import (
	"checkpoint2/internal/domain"
	"database/sql"
	"errors"
	"log"
)

type sqlStoreAppointment struct {
	db *sql.DB
}

func NewSQLStoreAppointment(db *sql.DB) AppointmentStoreInterface {
	return &sqlStoreAppointment{
		db: db,
	}
}

func (s *sqlStoreAppointment) ReadById(id int) (domain.Appointment, error) {
	queryGetById := `SELECT appointment.id, patient.id, patient.surname, patient.name, patient.rg, patient.registration_date, 
	                dentist.id, dentist.surname, dentist.name, dentist.registration, 
					appointment.date, appointment.description 
					FROM appointment 
					INNER JOIN patient 
					ON patient.id = appointment.patient_id 
					INNER JOIN dentist 
					ON dentist.id = appointment.dentist_id 
					WHERE appointment.id = ?`

	row := s.db.QueryRow(queryGetById, id)

	appointment := domain.Appointment{}

	err := row.Scan(
		&appointment.Id,
		&appointment.Patient.Id,
		&appointment.Patient.Surname,
		&appointment.Patient.Name,
		&appointment.Patient.RG,
		&appointment.Patient.RegistrationDate,
		&appointment.Dentist.Id,
		&appointment.Dentist.Surname,
		&appointment.Dentist.Name,
		&appointment.Dentist.Registration,
		&appointment.Date,
		&appointment.Description,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return appointment, errors.New("appointment not found")
	}

	if err != nil {
		return appointment, err
	}

	return appointment, nil
}

func (s *sqlStoreAppointment) ReadByRg(rg string) ([]domain.Appointment, error) {
	queryGetByRg := `SELECT appointment.id, patient.id, patient.surname, patient.name, patient.rg, patient.registration_date, 
					dentist.id, dentist.surname, dentist.name, dentist.registration, 
					appointment.date, appointment.description 
					FROM appointment 
					INNER JOIN patient 
					ON patient.id = appointment.patient_id 
					INNER JOIN dentist 
					ON dentist.id = appointment.dentist_id 
					WHERE patient.rg = ?`

	var appointments []domain.Appointment
	rows, err := s.db.Query(queryGetByRg, rg)
	if err != nil {
		return []domain.Appointment{}, err
	}

	defer rows.Close()

	for rows.Next() {
		appointment := domain.Appointment{}

		if err := rows.Scan(
			&appointment.Id,
			&appointment.Patient.Id,
			&appointment.Patient.Surname,
			&appointment.Patient.Name,
			&appointment.Patient.RG,
			&appointment.Patient.RegistrationDate,
			&appointment.Dentist.Id,
			&appointment.Dentist.Surname,
			&appointment.Dentist.Name,
			&appointment.Dentist.Registration,
			&appointment.Date,
			&appointment.Description,
		); err != nil {
			return appointments, err
		}
		appointments = append(appointments, appointment)
	}

	return appointments, nil
}

func (s *sqlStoreAppointment) CreateById(appointment domain.Appointment, idPatient int, idDentist int) (domain.Appointment, error) {
	queryInsert := "INSERT INTO appointment (patient_id, dentist_id, date, description) VALUES (?, ?, ?, ?)"

	stmt, err := s.db.Prepare(queryInsert)

	if err != nil {
		return domain.Appointment{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		idPatient,
		idDentist,
		appointment.Date,
		appointment.Description)
	if err != nil {
		return domain.Appointment{}, err
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return domain.Appointment{}, errors.New("failed to save appointment")
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return domain.Appointment{}, err
	}

	appointment.Id = int(lastId)

	appointment, err = s.ReadById(appointment.Id)
	if err != nil {
		return domain.Appointment{}, err
	}

	return appointment, nil
}

func (s *sqlStoreAppointment) CreateByRgAndRegistration(appointment domain.Appointment, rgPatient string, registrationDentist string) (domain.Appointment, error) {
	queryInsert := `INSERT INTO appointment (patient_id, dentist_id, date, description)
					VALUES ((SELECT patient.id FROM patient WHERE patient.rg = ?), 
					(SELECT dentist.id FROM dentist WHERE dentist.registration = ?), 
					?, ?)`

	stmt, err := s.db.Prepare(queryInsert)

	if err != nil {
		return domain.Appointment{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		rgPatient,
		registrationDentist,
		appointment.Date,
		appointment.Description)
	if err != nil {
		return domain.Appointment{}, err
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return domain.Appointment{}, errors.New("failed to save appointment")
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return domain.Appointment{}, err
	}

	appointment.Id = int(lastId)

	appointment, err = s.ReadById(appointment.Id)
	if err != nil {
		return domain.Appointment{}, err
	}

	return appointment, nil
}

func (s *sqlStoreAppointment) Update(id int, a domain.Appointment) (domain.Appointment, error) {
	queryUpdate  := "UPDATE appointment SET patient_id = ?, dentist_id = ?, date = ?, description = ? WHERE id = ?"

	persistedAppointment, err := s.ReadById(id)
	if err != nil {
		return domain.Appointment{}, errors.New("appointment not found")
	}

	persistedAppointment.Patient.Id = a.Patient.Id
	persistedAppointment.Dentist.Id = a.Dentist.Id
	persistedAppointment.Date = a.Date
	persistedAppointment.Description = a.Description

	result, err := s.db.Exec(
		queryUpdate,
		persistedAppointment.Patient.Id,
		persistedAppointment.Dentist.Id,
		persistedAppointment.Date,
		persistedAppointment.Description,
		id,
	)
	if err != nil {
		return domain.Appointment{}, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return domain.Appointment{}, err
	}
	log.Println(affectedRows)

	appointment, err := s.ReadById(id)
	if err != nil {
		return domain.Appointment{}, err
	}

	return appointment, nil
}

func (s *sqlStoreAppointment) Patch(id int, a domain.Appointment) (domain.Appointment, error) {
	queryUpdate  := "UPDATE appointment SET patient_id = ?, dentist_id = ?, date = ?, description = ? WHERE id = ?"

	appointment, err := s.ReadById(id)
	if err != nil {
		return domain.Appointment{}, errors.New("appointment not found")
	}

	if a.Patient.Id != 0 {
		appointment.Patient.Id = a.Patient.Id
	}
	if a.Dentist.Id != 0 {
		appointment.Dentist.Id = a.Dentist.Id
	}
	if a.Date != "" {
		appointment.Date = a.Date
	}
	if a.Description != "" {
		appointment.Description = a.Description
	}

	result, err := s.db.Exec(
		queryUpdate,
		appointment.Patient.Id,
		appointment.Dentist.Id,
		appointment.Date,
		appointment.Description,
		id,
	)
	if err != nil {
		return domain.Appointment{}, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return domain.Appointment{}, err
	}
	log.Println(affectedRows)

	appointment, err = s.ReadById(id)
	if err != nil {
		return domain.Appointment{}, err
	}

	return appointment, nil
}

func (s *sqlStoreAppointment) Delete(id int) error {
	queryDelete := "DELETE FROM appointment WHERE id = ?"

	result, err := s.db.Exec(queryDelete, id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	if affectedRows == 0 {
		return errors.New("appointment not found")
	}

	if err != nil {
		return err
	}

	return nil
}