package server

import (
	"fmt"
	"log"
	"loier/partilha/controllers"
	"net/http"
)

func Listen() {
	http.HandleFunc("/", health)
	http.HandleFunc("POST /partilha", controllers.PartilhaController)
	fmt.Println("Server is listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Health check!")
}
