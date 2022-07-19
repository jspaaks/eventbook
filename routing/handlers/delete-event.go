package handlers

import (
	"encoding/json"
	"net/http"
)

func DeleteEvent(events *[]Event) func(http.ResponseWriter, *http.Request) {

	handler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var encoder = json.NewEncoder(writer)
		var iEvent int64 = GetEventIndex(request)
		var nEvents int64 = GetNumberOfEvents(events)

		if iEvent <= nEvents {
			*events = append((*events)[:iEvent], (*events)[iEvent+1:]...)
			encoder.Encode(*events)
		}
	}
	return handler
}
