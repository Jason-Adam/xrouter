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

func New() XRouter {
	return &xRouter{}
}

type xRouter struct {
}

func (*xRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func (*xRouter) Use(middlewares ...func(http.Handler) http.Handler) {
}

func (*xRouter) Method(method string, path string, handlerFunc http.HandlerFunc) {
}

func (*xRouter) NotFound(handlerFunc http.HandlerFunc) {
}

func (*xRouter) MethodNotFound(handlerFunc http.HandlerFunc) {
}
