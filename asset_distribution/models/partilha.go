package models

type Partilha struct {
	Bem       Bem
	Herdeiros []Herdeiro
}

type PartilhaRequest struct {
	Bens     []Bem    `json:"bens"`
	Falecido Falecido `json:"falecido"`
}
