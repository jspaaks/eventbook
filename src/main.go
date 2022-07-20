package main

import (
	"log"
	"net/http"

	routing "github.com/jspaaks/eventbook/src/routing"
	types "github.com/jspaaks/eventbook/src/types"
)

func main() {

	var events []types.Event = []types.Event{
		{
			Text:     "the first text",
			MoreText: "More text pertaining to the first item",
		}, {
			Text:     "the second text",
			MoreText: "More text pertaining to the second item",
		},
	}

	var r = routing.CreateNewRouter().AddRoutesAndHandlers(&events)
	log.Fatal(http.ListenAndServe(":8080", r))
}
