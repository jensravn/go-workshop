package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
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
	dbPut(thing)
}

func dbPut(thing Thing) {
	bytes, err := json.Marshal(thing)
	if err != nil {
		log.Printf("err=%v", err)
	}
	os.WriteFile("thing.json", bytes, 0644)
}

func handleGetThing(w http.ResponseWriter, r *http.Request) {
	thing := dbGet()

	jsonThing, err := json.Marshal(thing)
	if err != nil {
		log.Printf("marshal err=%v", err)
	}
	w.Write(jsonThing)
}

func dbGet() *Thing {
	bytes, err := os.ReadFile("thing.json")
	if err != nil {
		log.Printf("readFile err=%v", err)
	}
	var thing Thing
	err = json.Unmarshal(bytes, &thing)
	if err != nil {
		log.Printf("unmarshal err=%v", err)
	}
	return &thing
}
