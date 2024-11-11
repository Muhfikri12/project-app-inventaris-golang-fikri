package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/model"
	apiservice "github.com/Muhfikri12/project-app-inventaris-golang-fikri/service/api_service"
	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/utils/validation"
	"github.com/go-chi/chi/v5"
)

type ItemApiHandler struct {
	ItemService *apiservice.ItemService
}

func NewItemHandler(serve *apiservice.ItemService) *ItemApiHandler {
	return &ItemApiHandler{ItemService: serve}
}


func (ih *ItemApiHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		validation.RequiredValidate(w, "Failed to parse form data: ")
		return
	}

	name := r.FormValue("name")
	categoryIdStr := r.FormValue("category_id")
	priceStr := r.FormValue("price")
	transactionDateStr := r.FormValue("transaction_date")

	if name == "" || categoryIdStr == "" || priceStr == "" || transactionDateStr == "" {
		validation.RequiredValidate(w, "All fields are required")
		return
	}

	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		validation.RequiredValidate(w, "Invalid category_id format")
		return
	}

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		validation.RequiredValidate(w, "Invalid price format")
		return
	}

	transactionDate, err := time.Parse("2006-01-02", transactionDateStr)
	if err != nil {
		validation.RequiredValidate(w, "Invalid date format, use YYYY-MM-DD")
		return
	}

	var imagePath string
	file, handler, err := r.FormFile("image")
	if err == nil && handler != nil {
		defer file.Close()

		ext := filepath.Ext(handler.Filename)
		timestamp := time.Now().Unix()
		imageName := fmt.Sprintf("%s-%d%s", strings.TrimSuffix(handler.Filename, ext), timestamp, ext)

		imagePath = "uploads/" + imageName

		dst, err := os.Create(imagePath)
		if err != nil {
			validation.BadResponse(w, "Failed to save image: ", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = io.Copy(dst, file)
		if err != nil {
			validation.BadResponse(w, "Failed to copy image: ", http.StatusInternalServerError)
			return
		}
	}

	totalDaysUsage := int(time.Since(transactionDate).Hours() / 24)

	item := &model.Items{
		Name:            name,
		CategoryId:      categoryId,
		Image:           imagePath,
		Price:           price,
		TransactionDate: transactionDate,
		TotalDaysUsage:  totalDaysUsage,
		Depreciation: 10,
	}

	err = ih.ItemService.CreateItemService(item)
	if err != nil {
		validation.RequiredValidate(w, "Failed to create item: ")
		return
	}

	validation.CreateResponse(w, "Create item successfully", item)
}


func (ih *ItemApiHandler) Items(w http.ResponseWriter, r *http.Request) {

	limitStr := r.URL.Query().Get("limit")
	pageStr := r.URL.Query().Get("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1 
	}

	pagination := &model.Response{}

	items, err := ih.ItemService.ShowAllItems(page, limit, pagination)
	if len(*items) <= 0 {
		validation.BadResponse(w, "items not found" + err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		validation.BadResponse(w, "Error :" + err.Error(), http.StatusInternalServerError)
		return
	}

	response := model.Response {
		Status: true,
		Page: page,
		Limit: limit,
		TotalItems: pagination.TotalItems,
		TotalPages: pagination.TotalPages,
		Data: items,
	}

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
}

func (ih *ItemApiHandler) ItemByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))


	if err != nil {
		validation.BadResponse(w, "Invalid Input ", http.StatusBadRequest)
		return
	}

	item, err := ih.ItemService.GetItemByID(id)
	
	if err != nil {
		validation.BadResponse(w, "item not found", http.StatusNotFound)
		return
	}

	totalDaysUsage := int(time.Since(item.TransactionDate).Hours() / 24) 

	item.TotalDaysUsage = totalDaysUsage

	validation.OKResponse(w, "Successfully", item)
}

func (ih *ItemApiHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		validation.BadResponse(w, "Invalid Input: ", http.StatusBadRequest)
		return
	}

	err = r.ParseMultipartForm(10 << 20)
	if err != nil {
		validation.BadResponse(w, "Error parsing form-data: ", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	categoryIdStr := r.FormValue("category_id")
	priceStr := r.FormValue("price")
	transactionDateStr := r.FormValue("transaction_date")

	if name == "" || categoryIdStr == "" || priceStr == "" || transactionDateStr == "" {
		validation.BadResponse(w, "All fields are required", http.StatusBadRequest)
		return
	}

	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		validation.BadResponse(w, "Invalid category_id: ", http.StatusBadRequest)
		return
	}

	price, err := strconv.Atoi(priceStr)
	if err != nil {
		validation.BadResponse(w, "Invalid price: ", http.StatusBadRequest)
		return
	}

	transactionDate, err := time.Parse("2006-01-02", transactionDateStr)
	if err != nil {
		validation.BadResponse(w, "Invalid transaction_date format: ", http.StatusBadRequest)
		return
	}

	var imagePath string
	file, handler, err := r.FormFile("image")
	if err == nil && handler != nil {
		defer file.Close()
		imagePath = "uploads/" + handler.Filename
		dst, err := os.Create(imagePath)
		if err != nil {
			validation.BadResponse(w, "Failed to save image: ", http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		_, err = io.Copy(dst, file)
		if err != nil {
			validation.BadResponse(w, "Failed to copy image: ", http.StatusInternalServerError)
			return
		}
	}

	item := &model.Items{
		ID: id,
		Name:            name,
		CategoryId:      categoryId,
		Image:           imagePath,
		Price:           price,
		TransactionDate: transactionDate,
	}

	err = ih.ItemService.UpdateItemID(id, item)
	if err != nil {
		validation.BadResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	totalDaysUsage := int(time.Since(item.TransactionDate).Hours() / 24)
	item.TotalDaysUsage = totalDaysUsage

	validation.CreateResponse(w, "Item successfully updated", item)
}

func (ih *ItemApiHandler) SoftDeleteItemHandler(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		validation.BadResponse(w, "Invalid Item: ", http.StatusBadRequest)
		return
	}

    err = ih.ItemService.SoftDeleteItemService(id)

    if err != nil {
        validation.BadResponse(w, "Failed to delete item: ", http.StatusNotFound)
        return
    }

    validation.OKResponse(w, "Item successfully deleted", nil)
}