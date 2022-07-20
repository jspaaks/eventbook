package routing

import (
	"github.com/gorilla/mux"
	handlers "github.com/jspaaks/eventbook/src/handlers"
	types "github.com/jspaaks/eventbook/src/types"
)

// make my own mux-based router using composition
type Router struct {
	*mux.Router
}

func CreateNewRouter() *Router {
	return &Router{mux.NewRouter()}
}

func (router *Router) AddRoutesAndHandlers(events *[]types.Event) *Router {
	router.HandleFunc("/api/events", handlers.AppendOne(events)).Methods("POST")
	router.HandleFunc("/api/events", handlers.GetAll(events)).Methods("GET")
	router.HandleFunc("/api/event/{index:[0-9]+}", handlers.GetOne(events)).Methods("GET")
	router.HandleFunc("/api/event/{index:[0-9]+}", handlers.UpdateOne(events)).Methods("PUT")
	router.HandleFunc("/api/event/{index:[0-9]+}", handlers.DeleteOne(events)).Methods("DELETE")
	return router
}
