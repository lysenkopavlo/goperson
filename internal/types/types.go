package types

import (
	"time"
)

type PostPerson struct {
	Name       string `json:"name"`
	Patronymic string `json:"patronymic,omitempty"`
	Surname    string `json:"surname"`
	Gender     string `json:"gender"`
	Age        int    `json:"age"`
	Country
}

type Country struct {
	CountryID   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}
type Person struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Patronymic string `json:"patronymic,omitempty"`
	Surname    string `json:"surname"`
	Gender     string `json:"gender"`
	Age        int    `json:"age"`
	CountryID  string `json:"country_id"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewPerson(age int, name, patronymic, surname, gender, countryID string) (Person, error) {
	return Person{
		Name:       name,
		Patronymic: patronymic,
		Surname:    surname,
		Gender:     gender,
		Age:        age,
		CountryID:  countryID,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}
