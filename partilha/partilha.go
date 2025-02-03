package partilha

type Pessoa struct {
	Nome          string
	DataObito     string
	DataCasamento string
	Conjuge       *Pessoa
	Filhos        []Pessoa
	Meeiro        bool
	Herdeiro      bool
}

type Bem struct {
	Nome  string
	Valor float64
}

type Partilha struct {
	Bem          Bem
	Distribuicao []Distribuicao
}

type Distribuicao struct {
	Pessoa     Pessoa
	Percentual float64
}

func (falecido Pessoa) CriarPartilha(listaDeBens []Bem) []Partilha {
	var partilhas []Partilha
	for _, bem := range listaDeBens {
		distribuicoes := calcularDistribuicoes(falecido)
		partilha := Partilha{
			Bem:          bem,
			Distribuicao: distribuicoes,
		}
		partilhas = append(partilhas, partilha)
	}
	return partilhas
}

func calcularDistribuicoes(falecido Pessoa) []Distribuicao {
	var distribuicoes []Distribuicao
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

func criarMeeiro(conjuge Pessoa, percentual float64, distribuicoes []Distribuicao) (float64, []Distribuicao) {
	percentual = percentual / 2
	distribuicao := Distribuicao{
		Pessoa:     conjuge,
		Percentual: percentual,
	}
	distribuicoes = append(distribuicoes, distribuicao)
	return percentual, distribuicoes
}

func criarHerdeiro(herdeiro Pessoa, quantidadeHerdeiros int, percentual float64) Distribuicao {
	distribuicao := Distribuicao{
		Pessoa:     herdeiro,
		Percentual: percentual / float64(quantidadeHerdeiros),
	}
	return distribuicao
}
