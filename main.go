package main

import (
	"context"
	"fmt"
	"loier/server"

	"github.com/go-pdf/fpdf"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	criarClientRedis()
	server.StartServer()
}

func criarClientRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-server:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Erro ao conectar ao Redis:", err)
		return
	}
	fmt.Println("Conexão bem sucedida ao Redis:", pong)
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
