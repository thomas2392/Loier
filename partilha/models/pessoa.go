package models

import "encoding/json"

type Familiar interface {
	NomeFamiliar() string
}

type Pessoa struct {
	NomeFalecido  string     `json:"nomeFalecido"`
	DataObito     string     `json:"dataObito"`
	DataCasamento string     `json:"dataCasamento"`
	Conjuge       Familiar   `json:"conjuge"`
	Filhos        []Familiar `json:"filhos"`
}

type rawPessoa struct {
	NomeFalecido  string            `json:"nomeFalecido"`
	DataObito     string            `json:"dataObito"`
	DataCasamento string            `json:"dataCasamento"`
	Conjuge       json.RawMessage   `json:"conjuge"`
	Filhos        []json.RawMessage `json:"filhos"`
}

type Conjuge struct {
	Nome   string `json:"nome"`
	Meeiro bool   `json:"meeiro"`
}

type Filho struct {
	Nome     string `json:"nome"`
	Herdeiro bool   `json:"herdeiro"`
}

func (c Conjuge) NomeFamiliar() string {
	return c.Nome
}

func (f Filho) NomeFamiliar() string {
	return f.Nome
}

func (p *Pessoa) UnmarshalJSON(data []byte) error {
	var raw rawPessoa
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw.Conjuge) != 0 {
		var conjuge Conjuge
		if err := json.Unmarshal(raw.Conjuge, &conjuge); err != nil {
			return err
		}
		p.Conjuge = conjuge
	}

	for _, filhoData := range raw.Filhos {
		var filho Filho
		if err := json.Unmarshal(filhoData, &filho); err != nil {
			return err
		}
		p.Filhos = append(p.Filhos, filho)
	}

	p.NomeFalecido = raw.NomeFalecido
	p.DataObito = raw.DataObito
	p.DataCasamento = raw.DataCasamento

	return nil
}
