package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type errorTrackRequest struct {
	ID     string   `json:"ID"`
	Name   string   `json:"name"`
	Image  string   `json:"image"`
	Age    int      `json:"age"`
	Errors []string `json:"errors"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/track/error/", trackError)

	log.Println("Starting server on :8080...")

	log.Fatal(http.ListenAndServe(":8080", mux))

}

func trackError(w http.ResponseWriter, r *http.Request) {
	var errorTrack errorTrackRequest

	err := decodeJSONBody(w, r, &errorTrack)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	js, err := json.Marshal(errorTrack)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
