package services

import (
	"loier/asset_distribution/models"
)

func CreateAssetDistribution(request models.AssetDistributionRequest) []models.AssetDistribution {
	var partilhas []models.AssetDistribution
	deceased := request.Deceased
	assets := request.Asssets
	for _, asset := range assets {
		herdeiros := calcularPercentuaisHeranca(deceased)
		partilha := models.AssetDistribution{
			Asset: asset,
			Heir:  herdeiros,
		}
		partilhas = append(partilhas, partilha)
	}
	return partilhas
}

func calcularPercentuaisHeranca(falecido models.Falecido) []models.Herdeiro {
	var herdeiros []models.Herdeiro
	percentual := 100.0

	if falecido.Conjuge != nil {
		conjuge, _ := falecido.Conjuge.(models.Conjuge)
		if conjuge.Meeiro {
			percentual, herdeiros = criarMeeiro(conjuge.RelativeName(), percentual, herdeiros)
		}
	}
	if len(falecido.Sons) < 1 {
		herdeiros[0].Percentual = 100.0
	} else {
		for _, son := range falecido.Sons {
			f, _ := son.(models.Son)
			herdeiro := criarHerdeiro(f.RelativeName(), len(falecido.Sons), percentual)
			herdeiros = append(herdeiros, herdeiro)
		}
	}
	return herdeiros
}

func criarMeeiro(nome string, percentual float64, herdeiros []models.Herdeiro) (float64, []models.Herdeiro) {
	percentual = percentual / 2
	herdeiro := models.Herdeiro{
		Nome:       nome,
		Percentual: percentual,
	}
	herdeiros = append(herdeiros, herdeiro)
	return percentual, herdeiros
}

func criarHerdeiro(nome string, quantidadeHerdeiros int, percentual float64) models.Herdeiro {
	herdeiro := models.Herdeiro{
		Nome:       nome,
		Percentual: percentual / float64(quantidadeHerdeiros),
	}
	return herdeiro
}
