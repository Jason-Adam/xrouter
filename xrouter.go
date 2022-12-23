package xrouter

import (
	"net/http"
)

type XRouter interface {
	http.Handler
	Use(middlewares ...func(http.Handler) http.Handler)
	Method(method string, path string, handlerFunc http.HandlerFunc)
	NotFound(handlerFunc http.HandlerFunc)
	MethodNotFound(handlerFunc http.HandlerFunc)
}

func New() *Mux {
	return &Mux{}
}

type Mux struct {
}

func (*Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func (*Mux) Use(middlewares ...func(http.Handler) http.Handler) {
}

func (*Mux) Method(method string, path string, handlerFunc http.HandlerFunc) {
}

func (*Mux) NotFound(handlerFunc http.HandlerFunc) {
}

func (*Mux) MethodNotFound(handlerFunc http.HandlerFunc) {
}
