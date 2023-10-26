package main

import "time"

type CreatePersonRequest struct {
	Name       string `json:"name"`
	Patronymic string `json:"patronymic,omitempty"`
	Surname    string `json:"surname"`
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

// type Country struct {
// 	CountryID   string  `json:"country_id"`
// 	Probability float32 `json:"probability"`
// }

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
