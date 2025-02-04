package models

type Pessoa struct {
	Nome          string
	DataObito     string
	DataCasamento string
	Conjuge       *Pessoa
	Filhos        []Pessoa
	Meeiro        bool
	Herdeiro      bool
}
