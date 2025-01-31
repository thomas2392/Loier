package server

import (
	"fmt"
	"log"
	"net/http"
)

func Listen() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
