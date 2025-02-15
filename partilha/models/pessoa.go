package models

type Familiar interface {
	NomeFamiliar() string
}

type Pessoa struct {
	NomeFalecido  string
	DataObito     string
	DataCasamento string
	Conjuge       Familiar
	Filhos        []Familiar
}

type Conjuge struct {
	Nome   string
	Meeiro bool
}

type Filho struct {
	Nome     string
	Herdeiro bool
}

func (c Conjuge) NomeFamiliar() string {
	return c.Nome
}

func (f Filho) NomeFamiliar() string {
	return f.Nome
}
