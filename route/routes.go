package route

import (
	"github.com/Muhfikri12/project-app-inventaris-golang-fikri/controller/api"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router *chi.Mux, itemHandler *api.ItemApiHandler, catHandler *api.CategoryHandler, investment *api.InvestmentHandler, auth *api.AuthHandler) {
	router.Post("/login", auth.LoginHandler)

	router.Route("/api", func(r chi.Router) {
		r.Use(auth.Middleware)

		r.Post("/items", itemHandler.CreateItem)
		r.Get("/items", itemHandler.Items)
		r.Get("/items/{id}", itemHandler.ItemByID)
		r.Put("/items/{id}", itemHandler.UpdateItem)
		r.Delete("/items/{id}", itemHandler.SoftDeleteItemHandler)

		r.Get("/categories", catHandler.Categories)
		r.Post("/category", catHandler.CreateCategory)
		r.Get("/category/{id}", catHandler.CategoryByID)
		r.Put("/category/{id}", catHandler.UpdateCategory)
		r.Delete("/category/{id}", catHandler.SoftDeleteCatHandler)

		r.Get("/investment", investment.GetInvestmentData)
		r.Get("/investment/{id}", investment.GetItemWithDepreciationID)
		r.Get("/investment/replaced", investment.GetReplacedItem)
	})
}
