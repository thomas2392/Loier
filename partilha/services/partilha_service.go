package services

import (
	"loier/partilha/models"
)

func CriarPartilha(falecido models.Pessoa, listaDeBens []models.Bem) []models.Partilha {
	var partilhas []models.Partilha
	for _, bem := range listaDeBens {
		distribuicoes := calcularDistribuicoes(falecido)
		partilha := models.Partilha{
			Bem:          bem,
			Distribuicao: distribuicoes,
		}
		partilhas = append(partilhas, partilha)
	}
	return partilhas
}

func calcularDistribuicoes(falecido models.Pessoa) []models.Distribuicao {
	var distribuicoes []models.Distribuicao
	percentual := 100.0

	if falecido.Conjuge.Meeiro {
		percentual, distribuicoes = criarMeeiro(*falecido.Conjuge, percentual, distribuicoes)
	}
	for _, filho := range falecido.Filhos {
		distribuicao := criarHerdeiro(filho, len(falecido.Filhos), percentual)
		distribuicoes = append(distribuicoes, distribuicao)
	}

	return distribuicoes
}

func criarMeeiro(conjuge models.Pessoa, percentual float64, distribuicoes []models.Distribuicao) (float64, []models.Distribuicao) {
	percentual = percentual / 2
	distribuicao := models.Distribuicao{
		Pessoa:     conjuge,
		Percentual: percentual,
	}
	distribuicoes = append(distribuicoes, distribuicao)
	return percentual, distribuicoes
}

func criarHerdeiro(herdeiro models.Pessoa, quantidadeHerdeiros int, percentual float64) models.Distribuicao {
	distribuicao := models.Distribuicao{
		Pessoa:     herdeiro,
		Percentual: percentual / float64(quantidadeHerdeiros),
	}
	return distribuicao
}
