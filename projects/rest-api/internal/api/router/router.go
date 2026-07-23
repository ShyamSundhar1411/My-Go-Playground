// Package router wires together HTTP routes and middleware for the API.
package router

import (
	"net/http"

	handler "github.com/ShyamSundhar1411/rest-api/internal/api/handlers"
	middleware "github.com/ShyamSundhar1411/rest-api/internal/api/middlewares"
)

func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handler.Health)

	return middleware.Logging(mux)
}
