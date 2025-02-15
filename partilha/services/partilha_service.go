package services

import (
	"loier/partilha/models"
)

func CriarPartilha(falecido models.Pessoa, listaDeBens []models.Bem) []models.Partilha {
	var partilhas []models.Partilha
	for _, bem := range listaDeBens {
		herdeiros := calcularPercentuaisHeranca(falecido)
		partilha := models.Partilha{
			Bem:       bem,
			Herdeiros: herdeiros,
		}
		partilhas = append(partilhas, partilha)
	}
	return partilhas
}

func calcularPercentuaisHeranca(falecido models.Pessoa) []models.Herdeiro {
	var herdeiros []models.Herdeiro
	percentual := 100.0

	if falecido.Conjuge != nil {
		conjuge, _ := falecido.Conjuge.(models.Conjuge)
		if conjuge.Meeiro {
			percentual, herdeiros = criarMeeiro(conjuge.NomeFamiliar(), percentual, herdeiros)
		}
	}
	for _, filho := range falecido.Filhos {
		f, _ := filho.(models.Filho)
		herdeiro := criarHerdeiro(f.NomeFamiliar(), len(falecido.Filhos), percentual)
		herdeiros = append(herdeiros, herdeiro)
	}

	return herdeiros
}

func criarMeeiro(nome string, percentual float64, herdeiros []models.Herdeiro) (float64, []models.Herdeiro) {
	percentual = percentual / 2
	herdeiro := models.Herdeiro{
		Nome:       "Teste",
		Percentual: percentual,
	}
	herdeiros = append(herdeiros, herdeiro)
	return percentual, herdeiros
}

func criarHerdeiro(nome string, quantidadeHerdeiros int, percentual float64) models.Herdeiro {
	herdeiro := models.Herdeiro{
		Nome:       nome,
		Percentual: percentual / float64(quantidadeHerdeiros),
	}
	return herdeiro
}
