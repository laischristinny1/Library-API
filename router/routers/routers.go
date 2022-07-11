package routers

import (
	"net/http"
	"github.com/gorilla/mux"
)
type Router struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
}

// Configure put all routers inside router
func Configure(r *mux.Router) *mux.Router {
	
	routers := BooksRouters

	for _, router := range routers{
		r.HandleFunc(router.URI, router.Function).Methods(router.Method)
	}

	return r
	
}