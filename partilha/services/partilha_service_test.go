package services

import (
	"loier/partilha/models"
	"testing"
)

func TestCriarMeeiroSemHerdeiros(t *testing.T) {
	percentual := 100.0
	percentualEsperado := percentual / 2
	var herdeiros []models.Herdeiro
	conjuge := models.Pessoa{
		Nome:          "ConjugeTeste",
		DataCasamento: "01/01/1990",
	}

	percentual, herdeiros = criarMeeiro(conjuge, percentual, herdeiros)

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
