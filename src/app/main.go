package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	// Internal packages
	"report"
)


type ajaxResponse struct {
	Ok bool
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/track/error/", trackError)

	log.Println("Starting server on :8080...")

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func trackError(w http.ResponseWriter, r *http.Request) {
	var visitReport report.VisitReport

	err := decodeJSONBody(w, r, &visitReport)

	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.msg, mr.status)
		} else {
			log.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	} else {
		go report.Create(visitReport)
	}

	w.Header().Set("Content-Type", "application/json")
	response := ajaxResponse{Ok: true}
	js, _ := json.Marshal(response)

	w.Write(js)

}
