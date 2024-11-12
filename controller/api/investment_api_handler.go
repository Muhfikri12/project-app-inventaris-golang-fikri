package api

import (
	"net/http"
	"strconv"

	apiservice "github.com/Muhfikri12/project-app-inventaris-golang-fikri/service/api_service"
	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/utils/validation"
	"github.com/go-chi/chi/v5"
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
		validation.BadResponse(w, "Data Not Found", http.StatusNotFound)
		return
	}

	validation.OKResponse(w, "Successfully Get data", investmentData)
}

func (h *InvestmentHandler) GetItemWithDepreciationID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		validation.RequiredValidate(w, "ID is not valid")
		return
	}

	item, err := h.service.GetItemWithDepreciationID(id)
	if err != nil {
		validation.BadResponse(w, "Data not found", http.StatusNotFound)
		return
	}

	validation.OKResponse(w, "Successflly", item)
}

func (h *InvestmentHandler) GetReplacedItem(w http.ResponseWriter, r *http.Request) {
	investmentData, err := h.service.GetIDShouldBeReplaced()
	if err != nil {
		validation.BadResponse(w, "Data Not Found", http.StatusNotFound)
		return
	}

	validation.OKResponse(w, "Successfully Get data", investmentData)
}