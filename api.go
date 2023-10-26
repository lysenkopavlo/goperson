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

	router.Get("/{id}", makeHTTPHandlerFunc(s.handleGetAccountByID))
	router.Post("/", makeHTTPHandlerFunc(s.handlePostAccount))
	router.Delete("/{id}", makeHTTPHandlerFunc(s.handleDeleteAccountByID))
	router.Put("/", makeHTTPHandlerFunc(s.handlePutAccount))

	log.Println("Server is up and listening on port", s.listenAddr)
	log.Fatal(http.ListenAndServe(s.listenAddr, router))
}

func (s *APIServer) handleGetAccountByID(rw http.ResponseWriter, req *http.Request) error {
	vars := chi.URLParam(req, "id")
	id, _ := strconv.Atoi(vars)
	// connecting to DB here
	p := Person{ID: id}
	return WriteJSON(rw, http.StatusOK, p)

}

func (s *APIServer) handlePostAccount(rw http.ResponseWriter, req *http.Request) error {
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
func (s *APIServer) handleDeleteAccountByID(rw http.ResponseWriter, req *http.Request) error {
	return nil
}

func (s *APIServer) handlePutAccount(rw http.ResponseWriter, req *http.Request) error {
	return nil
}
