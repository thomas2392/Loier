package services

import (
	"loier/asset_distribution/models"
)

func CreateAssetDistribution(request models.AssetDistributionRequest) []models.AssetDistribution {
	var assets_distribution []models.AssetDistribution
	deceased := request.Deceased
	assets := request.Asssets
	for _, asset := range assets {
		heirs := calculateInheritancePercentages(deceased)
		asset_distribution := models.AssetDistribution{
			Asset: asset,
			Heirs: heirs,
		}
		assets_distribution = append(assets_distribution, asset_distribution)
	}
	return assets_distribution
}

func calculateInheritancePercentages(deceased models.Falecido) []models.Heir {
	var heirs []models.Heir
	percentage := 100.0

	if deceased.Conjuge != nil {
		spouse, _ := deceased.Conjuge.(models.Conjuge)
		if spouse.Meeiro {
			percentage, heirs = criarMeeiro(spouse.RelativeName(), percentage, heirs)
		}
	}
	if len(deceased.Sons) < 1 {
		heirs[0].Percentage = 100.0
	} else {
		for _, son := range deceased.Sons {
			s, _ := son.(models.Son)
			herdeiro := criarHerdeiro(s.RelativeName(), len(deceased.Sons), percentage)
			heirs = append(heirs, herdeiro)
		}
	}
	return heirs
}

func criarMeeiro(nome string, percentual float64, herdeiros []models.Heir) (float64, []models.Heir) {
	percentual = percentual / 2
	herdeiro := models.Heir{
		Name:       nome,
		Percentage: percentual,
	}
	herdeiros = append(herdeiros, herdeiro)
	return percentual, herdeiros
}

func criarHerdeiro(nome string, quantidadeHerdeiros int, percentual float64) models.Heir {
	herdeiro := models.Heir{
		Name:       nome,
		Percentage: percentual / float64(quantidadeHerdeiros),
	}
	return herdeiro
}
