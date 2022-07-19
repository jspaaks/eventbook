package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetEventIndex(request *http.Request) int64 {
	var indexString string = mux.Vars(request)["index"]
	var iEvent, _ = strconv.ParseInt(indexString, 10, 0)
	return iEvent
}

func GetNumberOfEvents(events *[]Event) int64 {
	return int64(len(*events) - 1)
}
