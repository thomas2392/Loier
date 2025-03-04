package server

import (
	"loier/asset_distribution/controllers"
	"net/http"
)

func Router() {
	http.HandleFunc("GET /health", healthCheckHandler)
	http.HandleFunc("POST /asset_distribution", controllers.AssetDistributionController)
}
