package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")
var host = os.Getenv("HOST")

func StartServer() {
	Router()
	port_address := fmt.Sprintf(":%s", port)
	log.Printf("Server is listening on http://%s%s", host, port_address)
	if err := http.ListenAndServe(port_address, nil); err != nil {
		log.Printf("Error when starting server on http://%s%s.", host, port_address)
		log.Fatal(err)
	}
}
