package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Event struct {
	Text string `json:"text"`
}

var events []Event = []Event{
	{
		Text: "the first text",
	}, {
		Text: "the second text",
	},
}

func main() {
	var r *mux.Router = mux.NewRouter()
	r.HandleFunc("/api/events", createEvent).Methods("POST")
	r.HandleFunc("/api/events", readEvents).Methods("GET")
	r.HandleFunc("/api/event/{index:[0-9]+}", readEvent).Methods("GET")
	r.HandleFunc("/api/event/{index:[0-9]+}", updateEvent).Methods("PUT")
	r.HandleFunc("/api/event/{index:[0-9]+}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createEvent(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	var event Event
	_ = json.NewDecoder(request.Body).Decode(&event)
	events = append(events, event)
	json.NewEncoder(writer).Encode(events)
}

func readEvent(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var encoder = json.NewEncoder(writer)
	var iEvent int64 = getEventIndex(request)
	var nEvents int64 = getNumberOfEvents()
	if iEvent <= nEvents {
		encoder.Encode(events[iEvent])
	} else {
		encoder.Encode(&Event{})
	}
}

func readEvents(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(events)
}

func updateEvent(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var encoder = json.NewEncoder(writer)
	var iEvent int64 = getEventIndex(request)
	var nEvents int64 = getNumberOfEvents()
	if iEvent <= nEvents {
		var event Event
		_ = json.NewDecoder(request.Body).Decode(&event)
		events[iEvent] = event
		encoder.Encode(events)
	}
}

func deleteEvent(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var encoder = json.NewEncoder(writer)
	var iEvent int64 = getEventIndex(request)
	var nEvents int64 = getNumberOfEvents()

	if iEvent <= nEvents {
		events = append(events[:iEvent], events[iEvent+1:]...)
		encoder.Encode(events)
	}

}

func getEventIndex(request *http.Request) int64 {
	var indexString string = mux.Vars(request)["index"]
	var iEvent, _ = strconv.ParseInt(indexString, 10, 0)
	return iEvent
}

func getNumberOfEvents() int64 {
	return int64(len(events) - 1)
}
