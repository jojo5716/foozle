package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"foozle/pkg/report"
)

type errorTrackStruct struct {
	Project     string      `json:"project"`
	LoadedOn    string      `json:loadedOn`
	SessionTemp string      `json:sessionTemp`
	Session     string      `json:session`
	PageToken   string      `json:pageToken`
	EventTime   float32     `json:eventTime`
	Data        interface{} `json:"data"`
	Page        interface{} `json:page`
	Enviroment  interface{} `json:enviroment`
	MetaData    interface{} `json:metaData`
	Actions     interface{} `json:actions`
	UserInfo    interface{} `json:userInfo`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/track/error/", trackError)

	log.Println("Starting server on :8080...")

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func trackError(w http.ResponseWriter, r *http.Request) {
	var errorTrack errorTrackStruct
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
	} else {
		report.Create()
	}

	js, err := json.Marshal(errorTrack)

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
