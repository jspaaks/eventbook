package handlers

import (
	"encoding/json"
	"net/http"
)

func UpdateItem[T any](items *[]T) func(http.ResponseWriter, *http.Request) {
	handler := func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var encoder = json.NewEncoder(writer)
		var iItem int64 = GetItemIndex(request)
		var nItems int64 = GetNumberOfItems(items)
		if iItem <= nItems {
			var item T
			_ = json.NewDecoder(request.Body).Decode(&item)
			(*items)[iItem] = item
			encoder.Encode(items)
		}
	}
	return handler
}
