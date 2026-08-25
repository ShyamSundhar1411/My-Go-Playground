// Package router wires together HTTP routes and middleware for the API.
package router

import (
	"fmt"
	"net/http"
	"os"

	handler "github.com/ShyamSundhar1411/rest-api/internal/api/handlers"
	middleware "github.com/ShyamSundhar1411/rest-api/internal/api/middlewares"
	"github.com/swaggo/http-swagger"
	"github.com/ShyamSundhar1411/rest-api/docs"
)

func New() http.Handler {
	mux := http.NewServeMux()
	app_host := os.Getenv("APP_HOST")
	port := os.Getenv("PORT")
	if app_host == "" {
		app_host = "localhost"
	}
	host_url := fmt.Sprintf("%s:%s", app_host, port)
	docs.SwaggerInfo.Host = host_url
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	mux.HandleFunc("GET /healthz", handler.Health)
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), 
	))
	return middleware.Logging(mux)
}
