package main

import (
	"encoding/json"
	"net/http"
)

type thing struct {
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := thing{
			Message: "Hello world",
		}
		s, err := json.Marshal(&t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		b := []byte(s)
		w.Write(b)
	})
	http.ListenAndServe(":8080", nil)
}
