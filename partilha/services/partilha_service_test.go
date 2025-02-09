package services

import (
	"loier/partilha/models"
	"strconv"
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
	falecido := criarFalecido()
	falecido.Filhos = criarFilhos(2)
	herdeiros := calcularPercentuaisHeranca(falecido)
	totalPercentual := 0.0
	for _, h := range herdeiros {
		totalPercentual += h.Percentual
	}
	percentualEsperado := 100.0
	if totalPercentual != percentualEsperado {
		t.Fatalf("Resultado esperado era %v mas retornou %v", percentualEsperado, totalPercentual)
	}
}

func TestCalcularPercentualApenasConjuge(t *testing.T) {
	falecido := criarFalecido()
	falecido.Conjuge = criarConjuge()
	herdeiros := calcularPercentuaisHeranca(falecido)

	percentualTotal := 0.0
	for _, h := range herdeiros {
		percentualTotal += h.Percentual
	}

	percentualEsperado := 100.0

	if percentualTotal != percentualEsperado {
		t.Fatalf("O percentual esperado deveria ser %v mas foi %v", percentualEsperado, percentualTotal)
	}

}

func criarConjuge() *models.Pessoa {
	conjuge := models.Pessoa{
		Nome:          "ConjugeTeste",
		DataCasamento: "01/01/1990",
		Meeiro:        true,
	}

	return &conjuge
}

func criarFalecido() models.Pessoa {
	falcecido := models.Pessoa{
		Nome:          "FalcedidoTeste",
		DataObito:     "01/01/2020",
		DataCasamento: "01/01/2000",
		Conjuge:       criarConjuge(),
	}

	return falcecido
}

func criarFilhos(quantidade int) []models.Pessoa {
	var filhos []models.Pessoa
	for range quantidade {
		q := strconv.Itoa(quantidade)
		filho := models.Pessoa{
			Nome: "Filho" + q,
		}
		quantidade--
		filhos = append(filhos, filho)
	}

	return filhos
}
