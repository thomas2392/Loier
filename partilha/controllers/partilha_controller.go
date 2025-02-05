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
	partilhas := service.CriarPartilha(result, bens)

	for _, p := range partilhas {
		for _, distribuicao := range p.Distribuicao {
			fmt.Println("Bem:", p.Bem)
			fmt.Println("Herdeiro:", distribuicao.Pessoa.Nome)
			fmt.Println("Percentual:", distribuicao.Percentual)
		}
	}
}

func criarBens() []model.Bem {
	bem := model.Bem{
		Nome:  "Imovel",
		Valor: 100,
	}
	return []model.Bem{bem}
}
