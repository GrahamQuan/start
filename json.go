package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code >= 500 {
		log.Println("Responsing with 5** error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	responseWithJson(w, code, errResponse{
		Error: msg,
	})
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Failed to marshal json response: %v", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
