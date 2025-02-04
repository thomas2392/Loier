package main

import (
	"fmt"
	model "loier/partilha/models"
	service "loier/partilha/services"

	"github.com/go-pdf/fpdf"
)

func main() {
	//server.Listen()
	falecido := criarFalecido()
	bens := criarBens()
	partilhas := service.CriarPartilha(falecido, bens)

	for _, p := range partilhas {
		for _, distribuicao := range p.Distribuicao {
			fmt.Println("Bem:", p.Bem)
			fmt.Println("Herdeiro:", distribuicao.Pessoa.Nome)
			fmt.Println("Percentual:", distribuicao.Percentual)
		}
	}
}

func criarFalecido() model.Pessoa {
	falecido := model.Pessoa{
		Nome:          "Falecido",
		DataObito:     "03/02/2025",
		DataCasamento: "01/01/2025",
		Conjuge: &model.Pessoa{
			Nome:   "Conjuge",
			Meeiro: true,
		},
		Filhos: []model.Pessoa{
			{
				Nome: "Filho1",
			},
			{
				Nome: "Filho2",
			},
		},
	}
	return falecido
}

func criarBens() []model.Bem {
	bem := model.Bem{
		Nome:  "Imovel",
		Valor: 100,
	}
	return []model.Bem{bem}
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
