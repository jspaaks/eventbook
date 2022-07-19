package handlers

import (
	"encoding/json"
	"net/http"
)

func ReadEvent(events *[]Event) func(http.ResponseWriter, *http.Request) {
	handler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var encoder = json.NewEncoder(writer)
		var iEvent int64 = GetEventIndex(request)
		var nEvents int64 = GetNumberOfEvents(events)
		if iEvent <= nEvents {
			encoder.Encode((*events)[iEvent])
		} else {
			encoder.Encode(&Event{})
		}
	}
	return handler
}
