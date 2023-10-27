package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type APIServer struct {
	listenAddr string
	storage    DataBase
}

// APIerror is a type to hold server errors
type APIerror struct {
	Error string
}

// NewAPIserver return an instance of *APIserver with
// specified port address
func NewAPIserver(listenAddr string, typeOfStorage DataBase) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		storage:    typeOfStorage,
	}
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

func (s *APIServer) handleGetPersons(rw http.ResponseWriter, req *http.Request) error {

	persons, err := s.storage.GetPersons()
	if err != nil {
		return WriteJSON(rw, http.StatusBadRequest, fmt.Errorf("GetAllAccountFailed"))
	}

	return WriteJSON(rw, http.StatusOK, persons)

}

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

func (s *APIServer) handlePostPerson(rw http.ResponseWriter, req *http.Request) error {
	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusOK)

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}

	defer req.Body.Close()

	p := Person{}
	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Println(err)
	}
	persons := []Person{}
	persons = append(persons, p)

	fmt.Printf("%#v/n", persons)

	return nil
}
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

func (s *APIServer) handlePutPerson(rw http.ResponseWriter, req *http.Request) error {
	p, _ := NewPerson(18, "A", "B", "C", "M", "ua")

	// []Country{
	// 	{
	// 		CountryID: "ua",
	// 	}}}

	personsID, err := s.storage.AddPerson(p)
	if err != nil {
		return err
	}

	return WriteJSON(rw, http.StatusOK, fmt.Sprintf("Person added successfully. Id is %d", personsID))
}
