// Package router implements the router interface abstraction and server mux
package router

import "net/http"

type Router interface {
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}
