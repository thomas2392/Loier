package models

type AssetDistribution struct {
	Asset Bem
	Heirs []Heir
}

type AssetDistributionRequest struct {
	Asssets  []Bem    `json:"assets"`
	Deceased Falecido `json:"deceased"`
}
