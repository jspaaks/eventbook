package main

import (
	"log"
	"net/http"

	routing "github.com/jspaaks/eventbook/routing"
)

func main() {
	var r = routing.CreateNewRouter().AddRoutesAndHandlers()
	log.Fatal(http.ListenAndServe(":8080", r))
}
