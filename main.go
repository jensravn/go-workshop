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

func hello(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	data["msg"] = "hello"
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("marshal err=%v", err)
	}
	w.Write(jsonData)
}
