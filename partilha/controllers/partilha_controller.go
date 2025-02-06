package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	model "loier/partilha/models"
	service "loier/partilha/services"
)

func PartilhaController(w http.ResponseWriter, r *http.Request) {
	request, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var result model.Pessoa
	err = json.Unmarshal(request, &result)
	if err != nil {
		fmt.Println(err)
		return
	}
	bens := criarBens()
	response := service.CriarPartilha(result, bens)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func criarBens() []model.Bem {
	bem := model.Bem{
		Nome:  "Imovel",
		Valor: 100,
	}
	return []model.Bem{bem}
}
