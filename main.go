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
	r := chi.NewRouter()
	r.Get("/", handleGet)
	r.Put("/", handlePut)
	http.ListenAndServe(":8080", r)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile(thingTXT)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	var t thing
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s, err := json.Marshal(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b := []byte(s)
	os.WriteFile(thingTXT, b, 0644)
}
