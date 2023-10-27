package persons

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/lysenkopavlo/goperson/internal/types"
)

func UpdatePerson(update types.PostPerson) (types.Person, error) {
	p, err := types.NewPerson(
		update.Age,
		update.Name,
		update.Patronymic,
		update.Surname,
		update.Gender,
		update.CountryID,
	)
	if err != nil {
		return types.Person{}, err
	}

	return p, nil
}

func ExtractPostPerson(reqBody io.ReadCloser) (types.PostPerson, error) {
	person := types.PostPerson{}
	body, err := io.ReadAll(reqBody)
	if err != nil {
		return types.PostPerson{}, fmt.Errorf("%v", err)
	}

	defer reqBody.Close()

	err = json.Unmarshal(body, &person)

	if err != nil {
		return types.PostPerson{}, fmt.Errorf("%v", err)
	}
	return person, nil
}

func NewPostPerson(req *http.Request) (types.PostPerson, error) {

	pp, err := ExtractPostPerson(req.Body)

	if err != nil {
		return types.PostPerson{}, fmt.Errorf("%v", err)
	}

	return pp, nil
}

func UpdatePostPerson(link string, old types.PostPerson) (types.PostPerson, error) {

	resp, err := http.Get(link)
	if err != nil {
		return types.PostPerson{}, fmt.Errorf("%v", err)
	}

	updated, err := ExtractPostPerson(resp.Body)

	if err != nil {
		return types.PostPerson{}, fmt.Errorf("%v", err)

	}

	log.Println(old.Age, updated.Age)

	return updated, nil
}
