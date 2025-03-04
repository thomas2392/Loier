package models

import "encoding/json"

type Relative interface {
	RelativeName() string
}

type Falecido struct {
	NomeFalecido  string     `json:"nomeFalecido"`
	DataObito     string     `json:"dataObito"`
	DataCasamento string     `json:"dataCasamento"`
	Conjuge       Relative   `json:"conjuge"`
	Sons          []Relative `json:"sons"`
}

type rawPessoa struct {
	NomeFalecido  string            `json:"nomeFalecido"`
	DataObito     string            `json:"dataObito"`
	DataCasamento string            `json:"dataCasamento"`
	Conjuge       json.RawMessage   `json:"conjuge"`
	Sons          []json.RawMessage `json:"sons"`
}

type Conjuge struct {
	Nome   string `json:"nome"`
	Meeiro bool   `json:"meeiro"`
}

type Son struct {
	Name string `json:"name"`
	Heir bool   `json:"heir"`
}

func (c Conjuge) RelativeName() string {
	return c.Nome
}

func (s Son) RelativeName() string {
	return s.Name
}

func (f *Falecido) UnmarshalJSON(data []byte) error {
	var raw rawPessoa
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	if len(raw.Conjuge) != 0 {
		var conjuge Conjuge
		if err := json.Unmarshal(raw.Conjuge, &conjuge); err != nil {
			return err
		}
		f.Conjuge = conjuge
	}

	for _, sonData := range raw.Sons {
		var son Son
		if err := json.Unmarshal(sonData, &son); err != nil {
			return err
		}
		f.Sons = append(f.Sons, son)
	}

	f.NomeFalecido = raw.NomeFalecido
	f.DataObito = raw.DataObito
	f.DataCasamento = raw.DataCasamento

	return nil
}
