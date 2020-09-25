package Router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	fmt.Println("Route Init")
	register("GET", "/api", Hello, nil)
	fmt.Printf("%+v", routes)
}

func NewRouter() http.Handler {
	r := mux.NewRouter()
	for _, route := range routes {
		r.Methods(route.Method).
			Path(route.Pattern).
			Handler(route.Handler)
		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}
	handler := cors.Default().Handler(r)
	return handler
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Hello World")
	fmt.Fprintf(w, "Hello World")
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
