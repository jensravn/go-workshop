package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type thing struct {
	Message string `json:"message"`
}

const thingTXT = "thing.txt"

func main() {
	s := server{}
	s.routes()
	http.ListenAndServe(":8080", s.r)
}

type server struct {
	r *chi.Mux
}

func (s *server) routes() {
	r := chi.NewRouter()
	r.Get("/thing", s.handleGet)
	r.Put("/thing", s.handlePut)
	s.r = r
}

func (s *server) handleGet(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile(thingTXT)
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
	b, err := json.Marshal(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	os.WriteFile(thingTXT, b, 0644)
}
