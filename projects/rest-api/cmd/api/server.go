// Command api starts the REST API HTTP server.
package main

import (
	"log"
	"net/http"

	"github.com/ShyamSundhar1411/rest-api/internal/config"
	"github.com/ShyamSundhar1411/rest-api/internal/api/router"
)

func main() {
	cfg := config.Load()
	handler := router.New()

	addr := ":" + cfg.Port
	log.Printf("starting server in %s mode on %s", cfg.Env, addr)

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
