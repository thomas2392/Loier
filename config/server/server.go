package server

import (
	"log"
	"net/http"
)

func StartServer() {
	Router()
	log.Println("Server is listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("Error when starting server on http://localhost:8080.")
		log.Fatal(err)
	}
}
