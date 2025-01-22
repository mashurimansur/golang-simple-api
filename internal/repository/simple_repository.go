package repository

import (
	"database/sql"
	"golang-simple-api/internal/entity"
	"log"
)

type SimpleRepository struct {
	DB *sql.DB
}

func NewSimpleRepository(db *sql.DB) *SimpleRepository {
	return &SimpleRepository{
		DB: db,
	}
}

type SimpleRepositoryInterface interface {
	GetAll() ([]entity.Person, error)
	Create(person entity.Person) error
	GetByID(id int) (entity.Person, error)
}

func (s *SimpleRepository) GetAll() ([]entity.Person, error) {

	result, err := s.DB.Query("SELECT id, first_name, last_name, email, birth_date FROM person")

	if err != nil {
		return nil, err
	}

	defer result.Close()

	var persons []entity.Person

	for result.Next() {
		var person entity.Person

		err := result.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Email, &person.BirthDate)

		if err != nil {
			log.Println("Error scanning person: ", err.Error())
			return nil, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func (s *SimpleRepository) Create(person entity.Person) error {

	_, err := s.DB.Exec("INSERT INTO person (first_name, last_name, email, birth_date) VALUES (?,?,?,?)", person.FirstName, person.LastName, person.Email, person.BirthDate)

	if err != nil {
		log.Println("Error inserting person: ", err.Error())
		return err
	}

	return nil
}

func (s *SimpleRepository) GetByID(id int) (entity.Person, error) {

	result := s.DB.QueryRow("SELECT id, first_name, last_name, email, birth_date FROM person WHERE id = ?", id)

	var person entity.Person

	err := result.Scan(&person.ID, &person.FirstName, &person.LastName, &person.Email, &person.BirthDate)

	if err != nil {
		log.Println("Error scanning person: ", err.Error())
		return entity.Person{}, err
	}

	return person, nil
}
