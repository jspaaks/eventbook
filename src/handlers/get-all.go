package handlers

import (
	"encoding/json"
	"net/http"
)

func GetAll[T any](items *[]T) func(http.ResponseWriter, *http.Request) {
	handler := func(writer http.ResponseWriter, _ *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(*items)
	}
	return handler
}
