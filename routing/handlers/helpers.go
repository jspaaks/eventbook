package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetItemIndex(request *http.Request) int64 {
	var indexString string = mux.Vars(request)["index"]
	var iItem, _ = strconv.ParseInt(indexString, 10, 0)
	return iItem
}

func GetNumberOfItems[T any](items *[]T) int64 {
	return int64(len(*items) - 1)
}
