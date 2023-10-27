package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lysenkopavlo/goperson/internal/config"
	"github.com/lysenkopavlo/goperson/internal/persons"
	"github.com/lysenkopavlo/goperson/internal/repo/storage"
	"github.com/lysenkopavlo/goperson/internal/types"
)

// values to creating dummy struct in put method
const (
	age        = 18
	name       = "A"
	patronymic = "B"
	surname    = "C"
	gender     = "M"
	countryID  = "ua"
)

type APIServer struct {
	listenAddr     string
	storage        storage.DataBase
	ExternalSource config.ExternalLinks
}

// APIerror is a type to hold server errors
type APIerror struct {
	Error string
}

// NewAPIserver return an instance of *APIserver with
// specified port address
func NewAPIserver(listenAddr string, typeOfStorage storage.DataBase, externalLinks config.ExternalLinks) *APIServer {
	return &APIServer{
		listenAddr:     listenAddr,
		storage:        typeOfStorage,
		ExternalSource: externalLinks,
	}
}

// Run turns on a chi router
// and ListenAndServe function
func (s *APIServer) Run() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", makeHTTPHandlerFunc(s.handleGetPersons))
	router.Get("/{id}", makeHTTPHandlerFunc(s.handleGetPersonByID))
	router.Post("/", makeHTTPHandlerFunc(s.handlePostPerson))
	router.Delete("/{id}", makeHTTPHandlerFunc(s.handleDeletePersonByID))
	router.Put("/", makeHTTPHandlerFunc(s.handlePutPerson))

	log.Println("Server is up and listening on port", s.listenAddr)
	log.Fatal(http.ListenAndServe(s.listenAddr, router))
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

// apiFunc a custom HandleFunc to return an error
type apiFunc func(http.ResponseWriter, *http.Request) error

// makeHTTPHandlerFunc( a wrapper for custom type apiFunc
// to be of type HandlerFunc
func makeHTTPHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if err := f(rw, r); err != nil {
			WriteJSON(rw, http.StatusBadRequest, APIerror{Error: err.Error()})
		}
	}
}

// handleGetPersons returns all the persons from table
func (s *APIServer) handleGetPersons(rw http.ResponseWriter, req *http.Request) error {

	persons, err := s.storage.GetPersons()
	if err != nil {
		return WriteJSON(rw, http.StatusBadRequest, fmt.Errorf("GetAllAccountFailed"))
	}

	return WriteJSON(rw, http.StatusOK, persons)

}

// handleGetPersonByID finds person by given id
func (s *APIServer) handleGetPersonByID(rw http.ResponseWriter, req *http.Request) error {
	vars := chi.URLParam(req, "id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		return fmt.Errorf("error while converting string into integer: %v\n", err)
	}
	person, err := s.storage.GetPersonByID(id)
	if err != nil {
		return WriteJSON(rw, http.StatusNotFound, err)
	}

	return WriteJSON(rw, http.StatusOK, person)

}

// handlePostPerson handles POST request from random API
func (s *APIServer) handlePostPerson(rw http.ResponseWriter, req *http.Request) error {

	// Get a JSON struct from random API
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)

	// Create PostPerson struct to put received values
	receivedPerson, err := persons.NewPostPerson(req)
	if err != nil {
		return WriteJSON(rw, http.StatusBadRequest, err)
	}

	// Enrich struct fields with age
	upgraded, err := persons.EnrichPostPerson(s.ExternalSource.AgeLink, "age", receivedPerson)
	if err != nil {
		return WriteJSON(rw, http.StatusBadRequest, err)
	}

	upgraded, err = persons.EnrichPostPerson(s.ExternalSource.GenderLink, "gender", upgraded)

	if err != nil {
		return WriteJSON(rw, http.StatusBadRequest, err)
	}

	// update a Person struct before putting it into db
	upgraded, err = persons.EnrichPostPerson(s.ExternalSource.NationalityLink, "nationality", upgraded)
	if err != nil {
		return WriteJSON(rw, http.StatusBadRequest, err)
	}

	personToPut, err := persons.EnrichPerson(upgraded)
	if err != nil {
		return WriteJSON(rw, http.StatusBadRequest, err)

	}
	id, err := s.storage.AddPerson(personToPut)
	if err != nil {
		return WriteJSON(rw, http.StatusBadRequest, err)
	}

	return WriteJSON(rw, http.StatusOK, id)

}

// handleDeletePersonByID delete an person by given id
func (s *APIServer) handleDeletePersonByID(rw http.ResponseWriter, req *http.Request) error {
	vars := chi.URLParam(req, "id")
	id, err := strconv.Atoi(vars)
	if err != nil {
		return fmt.Errorf("error while converting string into integer: %v\n", err)
	}
	err = s.storage.DeletePersonByID(id)
	if err != nil {
		return WriteJSON(rw, http.StatusNotFound, err)
	}

	return WriteJSON(rw, http.StatusOK, fmt.Sprintf("Person with id %d deleted", id))
}

// handlePutPerson puts person into table
func (s *APIServer) handlePutPerson(rw http.ResponseWriter, req *http.Request) error {
	p, err := types.NewPerson(age, name, patronymic, surname, gender, countryID)

	if err != nil {
		return WriteJSON(rw, http.StatusNotFound, err)
	}

	personsID, err := s.storage.AddPerson(p)
	if err != nil {
		return err
	}

	return WriteJSON(rw, http.StatusOK, fmt.Sprintf("Person added successfully. Id is %d", personsID))
}
