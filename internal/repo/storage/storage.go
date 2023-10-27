package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/lysenkopavlo/goperson/internal/types"
)

// Declaring an interface where we us it
type DataBase interface {
	GetPersons() ([]types.Person, error)
	GetPersonByID(int) (types.Person, error)
	DeletePersonByID(int) error
	AddPerson(types.Person) (int64, error)
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

// GetPersons commands database to give all the persons
func (psql *PostgresRepo) GetPersons() ([]types.Person, error) {
	query := `SELECT * FROM persons`

	rows, err := psql.DB.Query(query)
	if err != nil {
		log.Printf("error while querying is: %v\n", err)
	}

	defer rows.Close()

	persons := []types.Person{}

	for rows.Next() {
		var p types.Person

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Patronymic,
			&p.Surname,
			&p.Age,
			&p.Gender,
			&p.CountryID,
			&p.CreatedAt,
			&p.UpdatedAt,
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

// GetPersonByID commands database to give one person with specified id
func (psql *PostgresRepo) GetPersonByID(id int) (types.Person, error) {
	p := types.Person{}
	query := `	SELECT * FROM persons WHERE id=$1`

	row := psql.DB.QueryRow(query, id)

	err := row.Scan(
		&p.ID,
		&p.Name,
		&p.Patronymic,
		&p.Surname,
		&p.Age,
		&p.Gender,
		&p.CountryID,
		&p.CreatedAt,
		&p.UpdatedAt,
	)

	if err != nil {
		return p, fmt.Errorf("Getting a person by ID failed: %v\n", err)
	}

	return p, nil
}

// DeletePersonByID commands database to delete person with specified id
func (psql *PostgresRepo) DeletePersonByID(id int) error {

	query := `DELETE FROM persons WHERE id=$1`

	_, err := psql.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Deleting a row failed: %v\n", err)
	}

	log.Printf("Person with id %d deleted", id)

	return nil
}

// AddPerson commands database to add one person
func (psql *PostgresRepo) AddPerson(p types.Person) (int64, error) {

	var personsID int
	query := `
	INSERT INTO persons
	(name, patronymic ,surname, gender, age, country_id, created_at, updated_at ) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	returning id`

	err := psql.DB.QueryRow(query, p.Name, p.Patronymic, p.Surname, p.Gender, p.Age, p.CountryID, p.CreatedAt, p.UpdatedAt).Scan(&personsID)

	if err != nil {
		return -1, fmt.Errorf("Inserting a row failed: %v\n", err)
	}

	if err != nil {
		return -1, fmt.Errorf("Inserting a row failed: %v\n", err)
	}

	log.Print("Person added successfully")
	return int64(personsID), nil
}
