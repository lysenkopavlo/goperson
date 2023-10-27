package persons

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"

	"github.com/lysenkopavlo/goperson/internal/types"
)

// NewPostPerson creates an example of received entity
func NewPostPerson(req *http.Request) (types.PostPerson, error) {

	pp, err := ExtractPostPerson(req.Body)

	if err != nil {
		return types.PostPerson{}, fmt.Errorf("error while extracting value of type PostPerson from body is: %v", err)
	}

	return pp, nil
}

// ExtractPostPerson unmarshalies given entity into PostPerson struct
func ExtractPostPerson(body io.ReadCloser) (types.PostPerson, error) {
	person := types.PostPerson{}
	person.Country = append(person.Country, types.Country{})

	bytes, err := io.ReadAll(body)
	if err != nil {
		return types.PostPerson{}, fmt.Errorf("error while response/request body reading %v", err)
	}

	defer body.Close()

	err = json.Unmarshal(bytes, &person)

	if err != nil {
		return types.PostPerson{}, fmt.Errorf("error while unmarshalling is: %v", err)
	}
	return person, nil
}

// EnrichPostPerson goes by given link and enriches PostPerson struct
func EnrichPostPerson(link, enrichment string, target types.PostPerson) (types.PostPerson, error) {

	resp, err := http.Get(link)
	if err != nil {
		return types.PostPerson{}, fmt.Errorf("error while requesting link %s is %v", link, err)
	}

	updated, err := ExtractPostPerson(resp.Body)

	if err != nil {
		return types.PostPerson{}, fmt.Errorf("error while extracting value of type PostPerson from body is: %v", err)

	}

	switch enrichment {
	case "age":
		target.Age = updated.Age
	case "gender":
		target.Gender = updated.Gender
	case "nationality":
		target.Country[0].CountryID = countryPicker(updated)
	}

	return target, nil
}

// EnrichPerson takes an updated PostPerson type and returns enriched Person type
// to put it into database
func EnrichPerson(update types.PostPerson) (types.Person, error) {
	tmp := make([]types.Country, 1)
	if len(update.Country) > 0 {
		tmp[0].CountryID = update.Country[0].CountryID
	}
	p, err := types.NewPerson(
		update.Age,
		update.Name,
		update.Patronymic,
		update.Surname,
		update.Gender,
		tmp[0].CountryID,
	)

	if err != nil {
		return types.Person{}, fmt.Errorf("error while enriching Person struct is: %v", err)
	}

	return p, nil
}

// countryPicker chooses country  with the hightest probability
func countryPicker(p types.PostPerson) string {
	if len(p.Country) == 0 {
		return "no country"
	}
	s := p.Country
	sort.Slice(s, func(i, j int) bool {
		return s[i].Probability > s[j].Probability
	})

	return s[0].CountryID
}
