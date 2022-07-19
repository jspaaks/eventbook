package routing

import (
	"github.com/gorilla/mux"
	handlers "github.com/jspaaks/eventbook/routing/handlers"
)

var events []handlers.Event = []handlers.Event{
	{
		Text: "the first text",
	}, {
		Text: "the second text",
	},
}

// make my own mux-based router using composition
type Router struct {
	*mux.Router
}

func CreateNewRouter() *Router {
	return &Router{mux.NewRouter()}
}

func (router *Router) AddRoutesAndHandlers() *Router {
	router.HandleFunc("/api/events", handlers.CreateEvent(&events)).Methods("POST")
	router.HandleFunc("/api/events", handlers.ReadEvents(&events)).Methods("GET")
	router.HandleFunc("/api/event/{index:[0-9]+}", handlers.ReadEvent(&events)).Methods("GET")
	router.HandleFunc("/api/event/{index:[0-9]+}", handlers.UpdateEvent(&events)).Methods("PUT")
	router.HandleFunc("/api/event/{index:[0-9]+}", handlers.DeleteEvent(&events)).Methods("DELETE")
	return router
}
