package server

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	Router()
	fmt.Println("Server is listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
