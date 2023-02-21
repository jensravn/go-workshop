package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:8080", nil)
}

type Thing struct {
	Msg string `json:"msg"`
}

func hello(w http.ResponseWriter, r *http.Request) {
	thing := Thing{
		Msg: "hello",
	}
	jsonThing, err := json.Marshal(thing)
	if err != nil {
		log.Printf("marshal err=%v", err)
	}
	w.Write(jsonThing)
}
