package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type errorTrackRequest struct {
	ID    string `json:"ID"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Age   int    `json:"age"`
}

func main() {

}

func New() {
	mux := http.NewServeMux()
	mux.HandleFunc("/track/error/", trackError)

	log.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}

func trackError(w http.ResponseWriter, r *http.Request) {
	var p errorTrackRequest

	err := decodeJSONBody(w, r, &p)
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

	fmt.Fprintf(w, "Person: %+v", p)
}
