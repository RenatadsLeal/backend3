package store

import (
	"checkpoint2/internal/domain"
	"database/sql"
	"errors"
	"log"
)

type sqlStoreDentist struct {
	db *sql.DB
}

func NewSQLStoreDentist(db *sql.DB) DentistStoreInterface {
	return &sqlStoreDentist{
		db: db,
	}
}

func (s *sqlStoreDentist) ReadById(id int) (domain.Dentist, error) {
	queryGetById := "SELECT id, surname, name, registration FROM dentist where id = ?"

	row := s.db.QueryRow(queryGetById, id)

	dentist := domain.Dentist{}

	err := row.Scan(
		&dentist.Id,
		&dentist.Surname,
		&dentist.Name,
		&dentist.Registration,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return dentist, errors.New("dentist not found")
	}

	if err != nil {
		return dentist, err
	}

	return dentist, nil
}

func (s *sqlStoreDentist) ReadAll() ([]domain.Dentist, error) {
	queryGetAll := "SELECT id, surname, name, registration FROM dentist"

	var dentists []domain.Dentist
	rows, err := s.db.Query(queryGetAll)
	if err != nil {
		return []domain.Dentist{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var dentist domain.Dentist

		if err := rows.Scan(
			&dentist.Id,
			&dentist.Surname,
			&dentist.Name,
			&dentist.Registration,
		); err != nil {
			return dentists, err
		}

		dentists = append(dentists, dentist)
	}
	return dentists, nil
}

func (s *sqlStoreDentist) ReadByRegistration(registration string) (domain.Dentist, error) {
	queryGetByRegistration := "SELECT id, surname, name, registration FROM dentist where registration = ?"

	row := s.db.QueryRow(queryGetByRegistration, registration)

	dentist := domain.Dentist{}

	err := row.Scan(
		&dentist.Id,
		&dentist.Surname,
		&dentist.Name,
		&dentist.Registration,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return dentist, errors.New("dentist not found")
	}

	if err != nil {
		return dentist, err
	}

	return dentist, nil
}

func (s *sqlStoreDentist) Create(dentist domain.Dentist) (domain.Dentist, error) {
	queryInsert := "INSERT INTO dentist (surname, name, registration) VALUES (?, ?, ?)"

	stmt, err := s.db.Prepare(queryInsert)

	if err != nil {
		return domain.Dentist{}, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		dentist.Surname,
		dentist.Name,
		dentist.Registration)
	if err != nil {
		return domain.Dentist{}, err
	}

	RowsAffected, _ := res.RowsAffected()
	if RowsAffected == 0 {
		return domain.Dentist{}, errors.New("failed to save")
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return domain.Dentist{}, err
	}

	dentist.Id = int(lastId)

	return dentist, nil
}

func (s *sqlStoreDentist) Update(id int, d domain.Dentist) (domain.Dentist, error) {
	queryUpdate  := "UPDATE dentist SET surname = ?, name = ?, registration = ? WHERE id = ?"

	dentist, err := s.ReadById(id)
	if err != nil {
		return domain.Dentist{}, errors.New("dentist not found")
	}

	dentist.Surname = d.Surname
	dentist.Name = d.Name
	dentist.Registration = d.Registration

	result, err := s.db.Exec(
		queryUpdate,
		dentist.Surname,
		dentist.Name,
		dentist.Registration,
		id,
	)
	if err != nil {
		return domain.Dentist{}, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}
	log.Println(affectedRows)

	return dentist, nil
}

func (s *sqlStoreDentist) Patch(id int, d domain.Dentist) (domain.Dentist, error) {
	queryUpdate  := "UPDATE dentist SET surname = ?, name = ?, registration = ? WHERE id = ?"

	dentist, err := s.ReadById(id)
	if err != nil {
		return domain.Dentist{}, errors.New("dentist not found")
	}

	if d.Surname != "" {
		dentist.Surname = d.Surname
	}

	if d.Name != "" {
		dentist.Name = d.Name
	}
	
	if d.Registration != "" {
		dentist.Registration = d.Registration
	}

	result, err := s.db.Exec(
		queryUpdate,
		dentist.Surname,
		dentist.Name,
		dentist.Registration,
		id,
	)
	if err != nil {
		return domain.Dentist{}, err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return domain.Dentist{}, err
	}
	log.Println(affectedRows)

	return dentist, nil
}

func (s *sqlStoreDentist) Delete(id int) error {
	queryDelete := "DELETE FROM dentist WHERE id = ?"

	result, err := s.db.Exec(queryDelete, id)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()

	if affectedRows == 0 {
		return errors.New("dentist not found")
	}

	if err != nil {
		return err
	}

	return nil
}