package router


import (
	"Library-API/router/routers"

	"github.com/gorilla/mux"
)

// Create will return a router with the configured routes
func Create() *mux.Router {
	r:= mux.NewRouter()
	return routers.Configure(r)
}