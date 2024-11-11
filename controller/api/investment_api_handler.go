package api

import (
	"net/http"

	apiservice "github.com/Muhfikri12/project-app-inventaris-golang-fikri/service/api_service"
	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/utils/validation"
)

type InvestmentHandler struct {
	service *apiservice.InvestmentService
}

func NewInvestmentHandler(service *apiservice.InvestmentService) *InvestmentHandler {
	return &InvestmentHandler{service}
}

func (h *InvestmentHandler) GetInvestmentData(w http.ResponseWriter, r *http.Request) {
	investmentData, err := h.service.GetTotalInvestment()
	if err != nil {
		validation.BadResponse(w, "Error : " + err.Error(), http.StatusInternalServerError)
		return
	}

	validation.OKResponse(w, "Successfully Get data", investmentData)
}