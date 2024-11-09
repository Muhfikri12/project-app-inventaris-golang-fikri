package route

import (
	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/controller/api"
	"github.com/go-chi/chi/v5"
)

func ItemsRoute(router *chi.Mux, itemHandler *api.ItemApiHandler) {
	router.Post("/api/items", itemHandler.CreateItem)
	router.Get("/api/items", itemHandler.Items)
	router.Get("/api/items/{id}", itemHandler.ItemByID)
	router.Put("/api/items/{id}", itemHandler.UpdateItem)
	router.Delete("/api/items/{id}", itemHandler.SoftDeleteItemHandler)
}

func CategoryRoute(router *chi.Mux, catHandler *api.CategoryHandler) {
	router.Get("/api/categories", catHandler.Categories)
	router.Post("/api/category", catHandler.CreateCategory)
}
