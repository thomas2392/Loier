package server

import (
	"encoding/json"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"status": "UP"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
