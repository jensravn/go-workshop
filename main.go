package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleThing)
	http.ListenAndServe("localhost:8080", nil)
}

type Thing struct {
	Msg string `json:"msg"`
}

func handleThing(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		handleGetThing(w, r)
	}
	if r.Method == "POST" {
		handlePostThing(w, r)
	}
}

func handlePostThing(w http.ResponseWriter, r *http.Request) {
	var thing Thing                               // Declare
	err := json.NewDecoder(r.Body).Decode(&thing) // Assign value
	if err != nil {
		log.Printf("err=%v", err)
	}
	log.Printf("%#v", thing)
}

func handleGetThing(w http.ResponseWriter, r *http.Request) {
	thing := Thing{
		Msg: "hello",
	}
	jsonThing, err := json.Marshal(thing)
	if err != nil {
		log.Printf("marshal err=%v", err)
	}
	w.Write(jsonThing)
}
