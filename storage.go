package main

import (
	"database/sql"
	"fmt"
	"log"
)

// Declaring an interface where we us it
type DataBase interface {
	GetPersons() ([]Person, error)
	GetPersonByID(int) (Person, error)
	DeletePersonByID(int) error
	AddPerson(Person) (int64, error)
	UpdatePerson(Person, string) error
}

// PostgresRepo type will implement DataBase interface
type PostgresRepo struct {
	DB *sql.DB
}

// NewPostgresRepo returns  database
func NewPostgresRepo(conn *sql.DB) *PostgresRepo {
	return &PostgresRepo{
		DB: conn,
	}
}

func (psql *PostgresRepo) GetPersons() ([]Person, error) {
	query := `SELECT * FROM persons`

	rows, err := psql.DB.Query(query)
	if err != nil {
		log.Printf("error while querying is: %v\n", err)
	}
	defer rows.Close()

	persons := []Person{}

	for rows.Next() {
		var p Person
		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Patronymic,
			&p.Surname,
			&p.Gender,
			&p.CountryID,
			&p.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("error from GetAllRows is: %v\n", err)
		}

		persons = append(persons, p)

	}

	//good practice to check the error after scanning rows
	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	fmt.Println("Recorded data:", persons)

	return persons, nil

}

func (psql *PostgresRepo) GetPersonByID(id int) (Person, error) {
	return Person{}, nil
}

func (psql *PostgresRepo) DeletePersonByID(id int) error {
	return nil
}

func (psql *PostgresRepo) AddPerson(p Person) (int64, error) {

	var personsID int
	query := `INSERT INTO persons(name, surname, gender, age, country_id, created_at, updated_at ) VALUES ($1, $2, $3, $4, $5, $6, $7) returning id`

	err := psql.DB.QueryRow(query, p.Name, p.Surname, p.Gender, p.Age, p.CountryID, p.CreatedAt, p.UpdatedAt).Scan(&personsID)

	if err != nil {
		return -1, fmt.Errorf("Inserting a row failed: %v\n", err)
	}

	if err != nil {
		return -1, fmt.Errorf("Inserting a row failed: %v\n", err)
	}

	log.Print("Person added successfully")
	return int64(personsID), nil
}
func (psql *PostgresRepo) UpdatePerson(p Person, str string) error {
	return nil
}
