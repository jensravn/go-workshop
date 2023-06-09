package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type thing struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			t := thing{
				Message: "Hello world",
			}
			s, err := json.Marshal(&t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			b := []byte(s)
			w.Write(b)
			return
		}
		if r.Method == http.MethodPut {
			var t thing
			err := json.NewDecoder(r.Body).Decode(&t)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			log.Printf("got thing: %#v", t)
			return
		}
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	})
	http.ListenAndServe(":8080", nil)
}
