package store

import (
	"checkpoint2/internal/domain"
	"database/sql"
	"errors"
	"log"
)

type sqlStorePatient struct {
	db *sql.DB
}

func NewSQLStorePatient(db *sql.DB) PatientStoreInterface {
	return &sqlStorePatient{
		db: db,
	}
}

func (s *sqlStorePatient) ReadById(id int) (domain.Patient, error) {
	queryGetById := "SELECT id, surname, name, rg, registration_date FROM patient where id = ?"

	row := s.db.QueryRow(queryGetById, id)

	patient := domain.Patient{}

	err := row.Scan(
		&patient.Id,
		&patient.Surname,
		&patient.Name,
		&patient.RG,
		&patient.RegistrationDate,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return patient, errors.New("patient not found")
	}

	if err != nil {
		return patient, err
	}

	return patient, nil
}

func (s *sqlStorePatient) ReadAll() ([]domain.Patient, error) {

	queryGetAll := "SELECT id, surname, name, rg, registration_date FROM patient"

	var patients []domain.Patient
	rows, err := s.db.Query(queryGetAll)
	if err != nil {
		return []domain.Patient{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var patient domain.Patient

		if err := rows.Scan(
			&patient.Id,
			&patient.Surname,
			&patient.Name,
			&patient.RG,
			&patient.RegistrationDate,
		); err != nil {
			return patients, err
		}

		patients = append(patients, patient)
	}
	return patients, nil
}

func (s *sqlStorePatient) ReadByRg(rg string) (domain.Patient, error) {
	queryGetByRg := "SELECT id, surname, name, rg, registration_date FROM patient where rg = ?"

	row := s.db.QueryRow(queryGetByRg, rg)

	patient := domain.Patient{}

	err := row.Scan(
		&patient.Id,
		&patient.Surname,
		&patient.Name,
		&patient.RG,
		&patient.RegistrationDate,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return patient, errors.New("patient not found")
	}

	if err != nil {
		return patient, err
	}

	return patient, nil
}

func (s *sqlStorePatient) Create(patient domain.Patient) (domain.Patient, error) {
	queryInsert := "INSERT INTO patient (surname, name, rg, registration_date) VALUES (?, ?, ?, ?)"

	stmt, err := s.db.Prepare(queryInsert)

	if err != nil {
		return domain.Patient{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		patient.Name,
		patient.Surname,
		patient.RG,
		patient.RegistrationDate)
	if err != nil {
		return domain.Patient{}, err
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return domain.Patient{}, errors.New("failed to save patient")
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return domain.Patient{}, err
	}

	patient.Id = int(lastId)

	return patient, nil
}

func (s *sqlStorePatient) Update(id int, patient domain.Patient) (domain.Patient, error) {
	queryUpdate  := "UPDATE patient SET surname = ?, name = ?, rg = ?, registration_date = ? WHERE id = ?"

	persistedPatient, err := s.ReadById(id)
	if err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}

	persistedPatient.Surname = patient.Surname
	persistedPatient.Name = patient.Name
	persistedPatient.RG = patient.RG
	persistedPatient.RegistrationDate = patient.RegistrationDate

	result, err := s.db.Exec(
		queryUpdate,
		persistedPatient.Surname,
		persistedPatient.Name,
		persistedPatient.RG,
		persistedPatient.RegistrationDate,
		id,
	)
	if err != nil {
		return domain.Patient{}, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return domain.Patient{}, err
	}
	log.Println(affectedRows)

	return persistedPatient, nil
}

func (s *sqlStorePatient) Patch(id int, patient domain.Patient) (domain.Patient, error) {
	queryUpdate  := "UPDATE patient SET surname = ?, name = ?, rg = ?, registration_date = ? WHERE id = ?"

	persistedPatient, err := s.ReadById(id)
	if err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}

	if patient.Surname != "" {
		persistedPatient.Surname = patient.Surname
	}
	if patient.Name != "" {
		persistedPatient.Name = patient.Name
	}
	if patient.RG != "" {
		persistedPatient.RG = patient.RG
	}
	if patient.RegistrationDate != "" {
		persistedPatient.RegistrationDate = patient.RegistrationDate
	}

	result, err := s.db.Exec(
		queryUpdate,
		persistedPatient.Surname,
		persistedPatient.Name,
		persistedPatient.RG,
		persistedPatient.RegistrationDate,
		id,
	)
	if err != nil {
		return domain.Patient{}, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return domain.Patient{}, err
	}
	log.Println(affectedRows)

	return persistedPatient, nil
}

func (s *sqlStorePatient) Delete(id int) error {
	queryDelete := "DELETE FROM patient WHERE id = ?"

	result, err := s.db.Exec(queryDelete, id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	if affectedRows == 0 {
		return errors.New("patient not found")
	}

	if err != nil {
		return err
	}

	return nil
}