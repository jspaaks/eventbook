package handlers

import (
	"encoding/json"
	"net/http"
)

func CreateEvent(events *[]Event) func(http.ResponseWriter, *http.Request) {

	handler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		var event Event
		_ = json.NewDecoder(request.Body).Decode(&event)
		*events = append(*events, event)
		json.NewEncoder(writer).Encode(*events)
	}

	return handler
}
