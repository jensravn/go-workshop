package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type thing struct {
	Message string `json:"message"`
}

const thingTXT = "thing.txt"

func main() {
	s := server{
		db: &dbFile{},
	}
	s.routes()
	http.ListenAndServe(":8080", s.r)
}

type server struct {
	r  *chi.Mux
	db *dbFile
}

func (s *server) routes() {
	r := chi.NewRouter()
	r.Get("/thing", s.handleGet)
	r.Put("/thing", s.handlePut)
	s.r = r
}

func (s *server) handleGet(w http.ResponseWriter, r *http.Request) {
	t, err := s.db.Get()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	b, err := json.Marshal(&t)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func (s *server) handlePut(w http.ResponseWriter, r *http.Request) {
	var t thing
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = s.db.Put(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type dbFile struct{}

func (db *dbFile) Get() (*thing, error) {
	b, err := os.ReadFile(thingTXT)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	var t thing
	err = json.NewDecoder(bytes.NewReader(b)).Decode(&t)
	if err != nil {
		return nil, fmt.Errorf("json decode: %w", err)
	}
	return &t, nil
}

func (db *dbFile) Put(t *thing) error {
	b, err := json.Marshal(&t)
	if err != nil {
		return fmt.Errorf("json marshal: %w", err)
	}
	os.WriteFile(thingTXT, b, 0644)
	if err != nil {
		return fmt.Errorf("write file: %w", err)
	}
	return nil
}
