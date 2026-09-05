// Command api starts the REST API HTTP server.
package main

import (
	"log"
	"net/http"

	"github.com/ShyamSundhar1411/rest-api/internal/api/router"
	"github.com/ShyamSundhar1411/rest-api/internal/config"
	"github.com/joho/godotenv"
)
// @title           School Management System API
// @version         1.0
// @description     This is a sample production-ready Go microservice.


// @BasePath  
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load environment variables")
	}
	cfg := config.Load()
	
	handler := router.New()
	config.InitDB()

	addr := ":" + cfg.Port
	log.Printf("starting server in %s mode on %s", cfg.Env, addr)

	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
