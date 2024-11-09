package api

import (
	"encoding/json"
	"net/http"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/model"
	apiservice "github.com/Muhfikri12/project-app-inventaris-golang-fikri/service/api_service"
	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/utils/validation"
)

type CategoryHandler struct {
	CategoryService *apiservice.ServiceCategory
}

func NewCategoryHandler(service *apiservice.ServiceCategory) *CategoryHandler {
	return &CategoryHandler{CategoryService: service}
}

func (ch *CategoryHandler) Categories(w http.ResponseWriter, r *http.Request) {


	categories, err := ch.CategoryService.GetAllCategory()

	if err != nil {
		validation.BadResponse(w, "Error :" + err.Error(), http.StatusInternalServerError)
		return
	}

	validation.OKResponse(w, "Successfully Retrived Categories", categories)
	
}

func (ch *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var category model.Categories

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		validation.BadResponse(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if category.Name == "" || category.Description == "" {
		validation.BadResponse(w, "Name and description are required", http.StatusBadRequest)
		return
	}

	err := ch.CategoryService.CreateCategoryService(&category)
	if err != nil {
		validation.BadResponse(w, "Failed to create category: "+err.Error(), http.StatusInternalServerError)
		return
	}

	validation.CreateResponse(w, category)
}
