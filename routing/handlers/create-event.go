package handlers

import (
	"encoding/json"
	"net/http"
)

func CreateItem[T any](items *[]T) func(http.ResponseWriter, *http.Request) {

	handler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		var item T
		_ = json.NewDecoder(request.Body).Decode(&item)
		*items = append(*items, item)
		json.NewEncoder(writer).Encode(*items)
	}

	return handler
}
