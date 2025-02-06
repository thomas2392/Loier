package server

import (
	"loier/partilha/controllers"
	"net/http"
)

func Router() {
	http.HandleFunc("GET /health", healthCheckHandler)
	http.HandleFunc("POST /partilha", controllers.PartilhaController)
}
