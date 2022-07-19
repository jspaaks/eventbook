package handlers

import (
	"encoding/json"
	"net/http"
)

func UpdateEvent(events *[]Event) func(http.ResponseWriter, *http.Request) {
	handler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var encoder = json.NewEncoder(writer)
		var iEvent int64 = GetEventIndex(request)
		var nEvents int64 = GetNumberOfEvents(events)
		if iEvent <= nEvents {
			var event Event
			_ = json.NewDecoder(request.Body).Decode(&event)
			(*events)[iEvent] = event
			encoder.Encode(events)
		}
	}
	return handler
}
