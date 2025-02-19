package main

import (
	redis "loier/config/redis"
	"loier/config/server"

	"github.com/go-pdf/fpdf"
)

func main() {
	redis.CreateRedisClient()
	server.StartServer()
}

func pdf() {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello, world!")
	err := pdf.OutputFileAndClose("hello.pdf")
	if err != nil {
		panic(err)
	}
}
