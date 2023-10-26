package main

import (
	"database/sql"
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
	return nil, nil
}

func (psql *PostgresRepo) GetPersonByID(id int) (Person, error) {
	return Person{}, nil
}

func (psql *PostgresRepo) DeletePersonByID(id int) error {
	return nil
}

func (psql *PostgresRepo) AddPerson(p Person) (int64, error) {
	return 0, nil
}
func (psql *PostgresRepo) UpdatePerson(p Person, str string) error {
	return nil
}
