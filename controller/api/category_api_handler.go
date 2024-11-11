package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/model"
	apiservice "github.com/Muhfikri12/project-app-inventaris-golang-fikri/service/api_service"
	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/utils/validation"
	"github.com/go-chi/chi/v5"
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

	validation.CreateResponse(w, "Successfully create Category", category)
}

func (ch *CategoryHandler) CategoryByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		validation.BadResponse(w, "Invalid Input ", http.StatusBadRequest)
		return
	}

	category, err := ch.CategoryService.GetCategoryByID(id)
	
	if err != nil {
		validation.BadResponse(w, "category not found", http.StatusNotFound)
		return
	}

	validation.OKResponse(w, "Successfully", category)
}

func (ch *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		validation.BadResponse(w, "Invalid category ID", http.StatusBadRequest)
		return
	}

	var category model.Categories
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		validation.BadResponse(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if category.Name == "" || category.Description == "" {
		validation.BadResponse(w, "Name and description are required", http.StatusBadRequest)
		return
	}

	if err := ch.CategoryService.UpdateCategory(id, &category); err != nil {
		validation.BadResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	validation.OKResponse(w, "Category successfully updated", category)
}

func (ch *CategoryHandler) SoftDeleteCatHandler(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		validation.BadResponse(w, "Invalid Category: ", http.StatusBadRequest)
		return
	}

    err = ch.CategoryService.SoftDeleteCatService(id)

    if err != nil {
        validation.BadResponse(w, err.Error(), http.StatusNotFound)
        return
    }

    validation.OKResponse(w, "Item successfully deleted", nil)
}
