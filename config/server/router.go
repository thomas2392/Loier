package server

import (
	"fmt"
	"loier/partilha/controllers"
	"net/http"
	"os"
)

func Router() {
	http.HandleFunc("GET /health", healthCheckHandler)
	http.HandleFunc("POST /partilha", controllers.PartilhaController)
	http.HandleFunc("GET /crash", crashApp)
}

func crashApp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Saindo da aplicação com erro...")
	os.Exit(1)
}
