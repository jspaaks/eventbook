package handlers

import (
	"encoding/json"
	"net/http"
)

func DeleteOne[T any](items *[]T) func(http.ResponseWriter, *http.Request) {

	handler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var encoder = json.NewEncoder(writer)
		var iItem int64 = GetItemIndex(request)
		var nItems int64 = GetNumberOfItems(items)

		if iItem <= nItems {
			*items = append((*items)[:iItem], (*items)[iItem+1:]...)
			encoder.Encode(*items)
		}
	}
	return handler
}
