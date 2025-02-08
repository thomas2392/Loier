package services

import (
	"loier/partilha/models"
	"testing"
)

func TestCriarMeeiroSemHerdeiros(t *testing.T) {
	percentual := 100.0
	percentualEsperado := percentual / 2
	var herdeiros []models.Herdeiro
	conjuge := criarConjuge()

	percentual, herdeiros = criarMeeiro(*conjuge, percentual, herdeiros)

	if percentual != percentualEsperado {
		t.Fatalf("Percentual esperado era metade %v, recebeu %v", percentualEsperado, percentual)
	}
	total := percentual
	for _, h := range herdeiros {
		total += h.Percentual
	}

	if total != 100.0 {
		t.Fatalf("Total do percentual deveria totalizar em 100, resultado foi %v", total)
	}
}

func TestCalcularPercentuaisHerancaSemConjuge(t *testing.T) {
	percentualEsperado := 100.0
	falecido := models.Pessoa{
		Nome:          "Falceido",
		DataObito:     "01/01/1990",
		DataCasamento: "01/01/2010",
		Filhos: []models.Pessoa{
			{
				Nome: "Filho01",
			},
			{
				Nome: "Filho02",
			},
		},
	}
	herdeiros := calcularPercentuaisHeranca(falecido)
	totalPercentual := 0.0
	for _, h := range herdeiros {
		totalPercentual += h.Percentual
	}

	if totalPercentual != percentualEsperado {
		t.Fatalf("Resultado esperado era %v mas retornou %v", percentualEsperado, totalPercentual)
	}
}

func criarConjuge() *models.Pessoa {
	conjuge := models.Pessoa{
		Nome:          "ConjugeTeste",
		DataCasamento: "01/01/1990",
	}

	return &conjuge
}
