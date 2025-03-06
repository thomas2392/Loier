package models

type AssetDistribution struct {
	Asset Bem
	Heir  []Herdeiro
}

type AssetDistributionRequest struct {
	Asssets  []Bem    `json:"assets"`
	Deceased Falecido `json:"deceased"`
}
