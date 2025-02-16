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

	percentual, herdeiros = criarMeeiro(conjuge.NomeFamiliar(), percentual, herdeiros)

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

func criarConjuge() models.Conjuge {
	conjuge := models.Conjuge{
		Nome:   "ConjugeTeste",
		Meeiro: true,
	}

	return conjuge
}

func criarFalecido() models.Pessoa {
	falcecido := models.Pessoa{
		NomeFalecido:  "FalcedidoTeste",
		DataObito:     "01/01/2020",
		DataCasamento: "01/01/2000",
		Conjuge:       criarConjuge(),
	}

	return falcecido
}

func criarFilhos(quantidade int) []models.Familiar {
	var filhos []models.Familiar
	for range quantidade {
		q := strconv.Itoa(quantidade)
		filho := models.Filho{
			Nome: "Filho" + q,
		}
		quantidade--
		filhos = append(filhos, filho)
	}

	return filhos
}
