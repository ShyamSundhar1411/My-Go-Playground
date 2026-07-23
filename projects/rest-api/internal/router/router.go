// Package router wires together HTTP routes and middleware for the API.
package router

import (
	"net/http"

	"github.com/ShyamSundhar1411/rest-api/internal/handler"
	"github.com/ShyamSundhar1411/rest-api/internal/middleware"
)

// New builds the top-level HTTP handler for the API, with all routes
// and global middleware attached. Add new resource routes here as
// handlers are implemented (users, orders, etc.).
func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handler.Health)

	return middleware.Logging(mux)
}
