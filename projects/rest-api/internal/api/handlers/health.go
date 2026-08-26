// Package handler contains HTTP handlers (controllers) for the API.
// Each resource (users, orders, etc.) should get its own file here,
// keeping request/response parsing thin and delegating business logic
// to the service layer.
package handler

import (
	"encoding/json"
	"net/http"
)

// Health godoc
// @Summary Health check
// @Description Returns the health status of the API
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /healthz [get]
func Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
	})
}
