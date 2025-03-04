package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	model "loier/asset_distribution/models"
	service "loier/asset_distribution/services"
)

func AssetDistributionController(w http.ResponseWriter, r *http.Request) {
	request, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	var result model.PartilhaRequest
	err = json.Unmarshal(request, &result)
	if err != nil {
		fmt.Println(err)
		return
	}
	response := service.CriarPartilha(result.Falecido, result.Bens)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
